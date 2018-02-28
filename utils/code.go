package utils

var Codes = map[int]string{
	0: "SUCCESS",
	10000: "Server error",
	10001: "Parameter error",  //参数错误
	10002: "Token error",  // token 错误
	10003: "Token expired",  // token过期
	10004: "Please login",
	10005: "Login forbidden", // 禁止登录
	10006: "Illegal request",
	10007: "You have no permission",
	10008: "Enter a valid URL",
	10009: "Policy is null",
	10010: "Can not support this Format ",


	//software code
	10701: "Software  does not exist",
	10702: "Software  already exist",
	10703: "Software id error",
	10704: "Please upload software",
	10705: "Create Software object error",
	10706: "The file type of Android_TV-Platform must be .apk",
	10707: "The file type of Android-Platform must be .apk",
	10708: "The hash of file has lost,can not be published",
	10709: "The file type of Table-Platform must be .7z",
	10710: "The file type of Camera-Platform must be .yak",
	10711: "The file type of PC-Platform must be .7z",

}