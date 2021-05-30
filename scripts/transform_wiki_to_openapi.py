import os
import re
from ruamel.yaml import YAML
import pandas as pd
from io import StringIO
from urllib.parse import urlparse, parse_qs
import json

pd.options.display.max_columns = 7
pd.options.display.width = 200
example = """
# Acquire Market Statistics

* Request description: Acquire real-time market data
* Request type: GET
* Signature required: No
* Request Url: [https://api.coinex.com/v1/market/ticker?market=BCHBTC](https://api.coinex.com/v1/market/ticker?market=BCHBTC)
* Request parameter:

  | name | type | required | description |
  | :--- | :--- | :--- | :--- |
  | market | String | Yes | See<API invocation description·market> |

* Return value description:

  | name | type | description |
  | :--- | :--- | :--- |
  | date | String | server time when returning |
  | last | String | latest price |
  | buy | String | buy 1 |
  | buy_amount | String | buy 1 amount|
  | sell | String | sell 1 |
  | sell_amount | String | sell 1 amount|
  | open | String | 24H open price |
  | high | String | 24H highest price |
  | low | String | 24H lowest price |
  | vol | String | 24H volume |

* Example:

```
# Request
GET https://api.coinex.com/v1/market/ticker?market=bchbtc
# Response
{
  "code": 0,
  "data": {
    "date": 1513865441609,  # server time when returning
    "ticker": {
      "buy": "10.00",           # buy 1
      "buy_amount": "10.00",    # buy 1 amount
      "open": "10",             # highest price
      "high": "10",             # highest price
      "last": "10.00",          # latest price
      "low": "10",              # lowest price
      "sell": "10.00",          # sell 1
      "sell_amount": "0.78",    # sell 1 amount
      "vol": "110"              # 24H volume
    }
  },
  "message": "Ok"
}
```
# Acquire All Market Data

* Request description: acquire all market data
* Request type: GET
* Signature required: No
* Request Url:[https://api.coinex.com/v1/market/ticker/all](https://api.coinex.com/v1/market/ticker/all)
* Request parameter:
  None

* Return value description:

  | name | type | description |
  | :--- | :--- | :--- |
  | date | String | server time when returning |
  | buy | String | buy 1 |
  | buy_amount | String | buy 1 amount|
  | high | String | 24H highest price |
  | last | String | latest price |
  | low | String | 24H lowest price |
  | sell | String | sell 1 |
  | sell_amount | String | sell 1 amount|
  | vol | String | 24H volumn |

* Example:

```
# Request
GET https://api.coinex.com/v1/market/ticker/all
# Response
{
    "code": 0,
    "data": {
        "date": 1513865441609,
        "ticker": {
            "BCHBTC": {
                "buy": "0.222",
                "buy_amount": "0.1",
                "open": "0.2322211",
                "high": "0.2322211",
                "last": "0.2322211",
                "low": "0.222",
                "sell": "0.3522211",
                "sell_amount": "0.11",
                "vol": "2.01430624"
            },
        }
    },
    "message": "Ok"
}
```
"""
api_regex = re.compile(r'[0-9]+([a-z]+)\_api', re.IGNORECASE | re.MULTILINE)
operation_regex = re.compile(r'^\#\ (.*)', re.MULTILINE)
details_regex = re.compile(
    r'^\#\ (?P<operation_name>[A-Za-z \-_]+)\W+'
    r'Request description:\W{0,1}(?P<description>[\w\W]+?)\W+'
    r'Request type\:\W(?P<operation_type>[\w]+)\W+'
    r'(Signature required:\W(?P<signature_required>[\w]+)\W+)?'
    r'(Rate limit\W+(?P<rate_limit>.*)\W+)?'
    r'(Request Header:\W(?P<request_header>[\w\W\<\>\\]+)\W+)?'
    r'Request Url\:\[(?P<request_url>https\:\/\/[\w\W]*?)\]\(https://[\w\W]+?\)\W+'
    r'(Request parameter:\W+?(?P<request_parameters>[\| \w\n\:\-\.\<\>\_]+|None)\W+?(?P<request_parameter_notes>[\w\W]*?)?\W+)?'
    r'(Return value description:\W+(?P<response_details>\|[\w\W]+\|)?\W+)?'
    r'Example\W+Request\W+(?P<request_example>.*)\W+'
    r'(Request\.Body(?P<request_body_example>[\W\w]+)\W+)?'
    r'Response\W+(?P<response_example>\{[\W\w]+\})\W+')
api_endpoints = []
for root, dirs, files in os.walk("../tmp/coinex_exchange_api.wiki", topdown=False):
    if '.git' in root:
        continue
    api_group = api_regex.findall(root)
    if len(api_group) == 0 or api_group[0] == 'rest':
        continue
    group = []
    for file in files:
        with open(os.path.join(root, file), encoding='utf-8') as f:
            data = f.read().strip('\n').replace('\u0008', '').replace(u'\xb7', ' ')
        operations = list(
            filter(lambda x: x != 'Request' and x != 'Response' and x != 'Request.Body', operation_regex.findall(data)))
        if len(operations) == 1:
            d = [m.groupdict() for m in details_regex.finditer(data)]
            group.append(d)
            if len(d) == 0:
                raise Exception('error')
        else:
            sections = [data]
            for operation in operations[::-1]:
                section = sections.pop(0).split(f'{operation}', 1)
                sections = [section[0]] + [f'# {operation}' + section[1]] + sections
            sections = list(filter(lambda x: x != '', sections))
            for section in sections:
                if section == '# ':
                    continue
                d = [m.groupdict() for m in details_regex.finditer(section)]
                group.append(d)
    api_endpoints.append((api_group, group))
dfs = []
for group_name, endpoints in api_endpoints:
    for endpoint in endpoints:
        df = pd.DataFrame.from_dict(endpoint)
        df['group_name'] = group_name
        dfs.append(df)
        # create
df = pd.concat(dfs)


class dotdict(dict):
    """dot.notation access to dictionary attributes"""
    __getattr__ = dict.get
    __setattr__ = dict.__setitem__
    __delattr__ = dict.__delitem__


from ruamel.yaml.representer import LiteralScalarString

yaml = YAML()


def str_presenter(dumper, data):
    """For handling multiline strings"""
    if len(data) > 72:  # check for multiline string
        return dumper.represent_scalar("tag:yaml.org,2002:str", LiteralScalarString(data), style="|")
    return dumper.represent_scalar("tag:yaml.org,2002:str", data)


yaml.representer.add_representer(str, str_presenter)


@yaml.register_class
class Dict(dotdict):
    def get(self, key, default=lambda: Dict()):
        val = dict.get(self, key)
        if val is None and default is not None:
            self[key] = default()
        return dict.__getitem__(self, key)

    @classmethod
    def to_yaml(cls, representer, data):
        if data is None:
            return representer.represent_none()
        return representer.represent_dict(dict(data))

    def set(self, value):
        self = Dict(value)


base = yaml.load("""
openapi: 3.0.2
info:
  title: CoinEx API
  description: |
    Open and simple, CoinEx API makes sure that you can build your own trading tools to achieve a more effective trading strategy. CoinEx API is now available for these features:
  version: 2021-05-29
servers:
  - url: https://api.coinex.com/v1
    description: Coinex Production Server
""")
openapi = Dict(base)
paths = openapi.get('paths')
components = openapi.get('components')
components.get('parameters')['CX-ACCESS-SIGN'] = yaml.load("""
CX-ACCESS-SIGN:
  in: header
  name: authorization
  required: true
  description: |
    Signature is required for Account API and trading API related interfaces.
    The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string.
    No signature is required for market API related interfaces.
    Use 32-bit MD5 Algorithm Signature
    Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.
  schema:
    type: string
  x-go-default-value: "auto"
  x-go-default: true
""")['CX-ACCESS-SIGN']
components.get('schemas')['UnknownResponse'] = {
    'type'       : 'object',
    'description': 'Unknown Response'
}
for group_name, endpoints in api_endpoints:
    for endpoint in [dotdict(l2) for l1 in endpoints for l2 in l1]:
        path = endpoint.request_url.replace('https://api.coinex.com/v1', '').split('?')[0]
        methods = paths.get(path)

        method = endpoint.operation_type.strip().lower()
        if method in methods:
            continue
        dmethod = methods.get(method)
        dmethod.summary = endpoint.operation_name
        dmethod.description = endpoint.description.strip(
            '*').strip() + (f'\nRate Limit: {endpoint.rate_limit}' if endpoint.rate_limit else '')
        dmethod.tags = list(group_name)
        dmethod.operationId = ''.join(
            [x.title() for x in endpoint.operation_name.replace('-', ' ').split(' ')])
        parameters = dmethod.get('parameters', default=lambda: [])
        if endpoint.signature_required and endpoint.signature_required.strip() == 'Yes':
            parameters.append({
                '$ref': '#/components/parameters/CX-ACCESS-SIGN'
            })
        types_lookup = {
            'String' : ('string', {}),
            'Integer': ('integer', {'format': 'int64'}),
            'Long'   : ('number', {'format': 'double'}),
            'Array'  : ('array', lambda: {'items': Dict({'type': 'object'})}),
            'Object' : ('object', {}),
            'bool'   : ('boolean', {})
        }
        if endpoint.request_parameters:
            parameters_table = pd.read_table(
                StringIO(endpoint.request_parameters), sep="|", header=0, skipinitialspace=True
            ).dropna(axis=1, how='all').iloc[1:]
            parameters_table.columns = list(map(lambda x: x.strip(), parameters_table.columns))
            if endpoint.request_body_example is None:
                parsed_url = urlparse(endpoint.request_example.split(' ', 1)[1])
                query_dict = parse_qs(parsed_url.query)
                for index, row in parameters_table.iterrows():
                    _type, format = types_lookup[row.type.strip().replace('Interger', 'Integer')]
                    parameters.append({
                        'in'      : 'query',
                        'name'    : row.get('name').strip(),
                        'required': True if 'yes' in row.required.lower() else False,
                        'schema'  : {
                            'type': _type,
                            **format
                        },
                    })
                    if pd.notna(row.description):
                        parameters[-1]['description'] = row.description.strip()
                    example = query_dict.get(row.get('name').strip(), None)
                    if example:
                        parameters[-1]['example'] = example[0]
            else:
                dmethod.get('requestBody').get('content')['application/json'] = {
                    'schema': {'$ref': f'#/components/schemas/{dmethod.operationId}Request'},
                }
                components.get('schemas')[f'{dmethod.operationId}Request'] = Dict({
                    'type'   : 'object',
                    'example': LiteralScalarString(endpoint.request_body_example.strip().strip('#'))
                })
                requestBody = components.get('schemas')[f'{dmethod.operationId}Request'].get('properties')
                for index, row in parameters_table.iterrows():
                    key = row.get('name').strip()
                    _type, format = types_lookup[row.type.strip().replace('Interger', 'Integer')]
                    requestBody[key] = {
                        'type': _type,
                        **format
                    }
                    if pd.notna(row.required) and 'yes' in row.required.lower():
                        components.get('schemas')[f'{dmethod.operationId}Request'].get('required',
                                                                                       default=lambda: []).append(key)
                    if pd.notna(row.description):
                        requestBody[key]['description'] = row.description.strip()

        responses = dmethod.get('responses')
        responses['200'] = Dict({
            'description': 'response info',
        })
        if endpoint.response_details:
            key = f'{dmethod.operationId}Response'
            responses.get('200').get('content')['application/json'] = Dict({
                'schema': {
                    'type'      : 'object',
                    'properties': {
                        'code'   : {
                            'type'  : 'integer',
                            'format': 'int64'
                        },
                        'message': {
                            'type': 'string'
                        },
                        'data'   : {'$ref': f'#/components/schemas/{key}'}
                    }
                }
            })
            response_table = pd.read_table(
                StringIO(endpoint.response_details), sep="|", header=0, skipinitialspace=True
            ).dropna(axis=1, how='all').iloc[1:]
            response_table.columns = list(map(lambda x: x.strip(), response_table.columns))
            components.get('schemas')[key] = Dict({
                'type'       : 'object',
                'description': f'{dmethod.operationId} Response Value',
            })
            if endpoint.response_example:
                try:
                    components.get('schemas')[key]['example'] = json.loads(endpoint.response_example.strip().strip('#'))['data']
                except json.decoder.JSONDecodeError:
                    components.get('schemas')[key]['example'] = LiteralScalarString(
                        endpoint.response_example.strip().strip('#'))
            responseBody = components.get('schemas')[key].get('properties')
            for index, row in response_table.iterrows():
                if '[' in row.get('name'):
                    continue
                name = row.get('name').strip().replace('\\', '')
                _type, format2 = types_lookup[row.get('type', 'Object').strip().replace('Interger', 'Integer')]
                responseBody[name] = {
                    'type': _type,
                }
                try:
                    responseBody[name].update(format2)
                except TypeError:
                    responseBody[name].update(format2())
                if pd.notna(row.description):
                    responseBody[name]['description'] = row.description.strip()
        else:
            responses.get('200').get('content')['application/json'] = Dict({
                'schema': {
                    'type'      : 'object',
                    'properties': {
                        'code'   : {
                            'type'  : 'integer',
                            'format': 'int64'
                        },
                        'message': {
                            'type': 'string'
                        },
                        'data'   : {
                            '$ref': f'#/components/schemas/UnknownResponse'
                        }
                    }
                }
            })

with open('../api/coinex.yml', 'w', encoding='utf-8') as f:
    yaml.dump(openapi, f)
