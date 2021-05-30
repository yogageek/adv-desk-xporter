package logic

import "porter/util"

//gql query全部添加上多語言欄位  ok
func TestQuery() {
	//machine status OK
	func() {
		res := QueryMachineStatuses()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//group ok
	func() {
		res := QueryGroups()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//machines ok
	func() {
		res := QueryMachines()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//ParameterMappingCode ok
	func() {
		res := QueryParameterMappings()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//ProfileMachine ok
	func() {
		res := QueryProfileMachines()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	// Parameter ok
	func() {
		res := QueryParameters("TWFjaGluZQ.YD86DsGunQAGpEaI", "")
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()
}

//gql translate多語言欄位資料
func TestAdd() {

}
