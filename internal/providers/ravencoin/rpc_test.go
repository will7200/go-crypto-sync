package ravencoin

import (
	"testing"
)

func Test_readFromRPCNode(t *testing.T) {
	type args struct {
		args rpcNodeArgs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"valid", args{args: rpcNodeArgs{url: "http://localhost:8766", userPassword: "file:///d:/crypto/ravencoin/.cookie"}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := readFromRPCNode(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFromRPCNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("readFromRPCNode() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
