package booking_test

//import (
//	"context"
//	"os"
//	"testing"
//)
//
//var ctx = context.Background()
//
//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}

// func BenchmarkCheckAccess(b *testing.B) {
// 	var host = "stag3:27000"
// 	conn, _ := grpc.NewClient(host)
// 	client := api_role.NewRoleClient(conn)
// 		b.RunParallel(func(bp *testing.PB) {
// 			for bp.Next() {
// 				for _, testCase := range testCases {
// 					access, err := client.CheckAccess(ctx, &api_role.CheckAccessRequest{
// 						RoleCode: testCase.roleCode,
// 						Module:   testCase.module,
// 						Action:   testCase.action,
// 					})
// 					if testCase.expectError {
// 						if err == nil {
// 							b.Fatalf("FAIL: %s\n\tCheckAccess() roleCode:%v | module:%v | action:%v expected an error, got %v",
// 								testCase.description, testCase.roleCode, testCase.module, testCase.action, access)
// 						}
// 					} else {
// 						if err != nil {
// 							b.Fatalf("FAIL: %s\n\tCheckAccess() roleCode:%v | module:%v | action:%v returns unexpected error %s",
// 								testCase.description, testCase.roleCode, testCase.module, testCase.action, err)
// 						}
// 						if access.Data != testCase.expected {
// 							b.Fatalf("FAIL: %s\n\tCheckAccess() roleCode:%v | module:%v | action:%v expected %v, got %v",
// 								testCase.description, testCase.roleCode, testCase.module, testCase.action, testCase.expected, access)
// 						}
// 					}
// 				}
// 			}
// 		})
// }
