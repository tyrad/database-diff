(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-b496b1c6"],{"018a":function(e,t,o){},"750e":function(e,t,o){"use strict";o("018a")},7518:function(e,t,o){},"911b":function(e,t,o){"use strict";var r=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],staticStyle:{margin:"0 auto",width:"300px"}},[o("div",{staticStyle:{color:"red","margin-bottom":"10px"}},[e._v(" 鉴权出错，请登录后再尝试 ")]),o("div",{staticStyle:{display:"flex","align-items":"center","justify-items":"center"}},[o("img",{staticStyle:{padding:"5px"},attrs:{src:e.imageUrl,alt:""}}),o("el-button",{staticStyle:{padding:"5px"},attrs:{size:"mini"},on:{click:e.refReloadCodeImage}},[e._v("点击刷新验证码")])],1),o("el-input",{staticStyle:{margin:"5px",width:"200px"},attrs:{placeholder:"请输入验证码",clearable:""},model:{value:e.code,callback:function(t){e.code=t},expression:"code"}})],1)},a=[],i=(o("d3b7"),o("b5f9")),n={name:"Login",data:function(){return{loading:!1,imageUrl:"",code:""}},methods:{getImageUrl:function(){return this.imageUrl?window.location.origin+this.imageUrl:""},login:function(){var e=this;this.loading=!0,Object(i["h"])(this.code).then((function(){e.$emit("loginSuccess")})).finally((function(){e.loading=!1}))},refReloadCodeImage:function(){var e=this;this.loading=!0,Object(i["f"])().then((function(t){console.log(t),e.imageUrl=t.data.path})).finally((function(){e.loading=!1}))}}},s=n,l=o("2877"),c=Object(l["a"])(s,r,a,!1,null,"36c194cc",null);t["a"]=c.exports},b5f9:function(e,t,o){"use strict";o.d(t,"a",(function(){return s})),o.d(t,"f",(function(){return l})),o.d(t,"h",(function(){return c})),o.d(t,"g",(function(){return u})),o.d(t,"b",(function(){return d})),o.d(t,"c",(function(){return f})),o.d(t,"i",(function(){return m})),o.d(t,"e",(function(){return h})),o.d(t,"d",(function(){return p})),o.d(t,"j",(function(){return g}));var r=o("5530"),a=o("b775"),i=o("4328"),n=o.n(i);function s(e){return Object(a["a"])({url:"/checkAccountValid",method:"post",data:n.a.stringify(Object(r["a"])({},e))})}function l(){return Object(a["a"])({url:"/getCodeImage",method:"get"})}function c(e){return Object(a["a"])({url:"/loginWithCode",method:"get",params:{code:e}})}function u(e){return Object(a["a"])({url:"/getHospitalCode",method:"get",params:{name:e}})}function d(e){return Object(a["a"])({url:"/createYxAccountOrUpdate",method:"post",data:n.a.stringify(Object(r["a"])({},e))})}function f(e,t){return Object(a["a"])({url:"/deleteDoctor",method:"get",params:{doctorId:e,phoneNo:t}})}function m(e){return Object(a["a"])({url:"/resetPassword",method:"get",params:{doctorId:e}})}function h(){return Object(a["a"])({url:"/getAreaCode",method:"get"})}function p(e){return Object(a["a"])({url:"/fixedRoleToDoctor",method:"get",params:{phone:e}})}function g(e){return Object(a["a"])({url:"/uploadFile",method:"post",data:e,headers:{"Content-Type":"multipart/form-data;"}})}},c971:function(e,t,o){"use strict";o("7518")},ce6f:function(e,t,o){"use strict";o.r(t);var r=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{staticClass:"th-check"},[o("el-dialog",{attrs:{title:"登录",visible:!e.loginSuccess,width:"60%","before-close":e.handleClose}},[o("login",{ref:"loginFormRef",on:{loginSuccess:function(t){e.loginSuccess=!0}}}),o("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:function(t){e.loginSuccess=!0}}},[e._v("取 消")]),o("el-button",{attrs:{type:"primary"},on:{click:e.loginSubmit}},[e._v("确 定")])],1)],1),o("YixinAnalz",{on:{reLogin:e.reLogin}})],1)},a=[],i=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}]},[o("el-form",{ref:"elForm",attrs:{model:e.formData,rules:e.rules,size:"medium","label-width":"100px"}},[o("el-form-item",{attrs:{label:"手机号",prop:"phone"}},[o("el-input",{style:{width:"100%"},attrs:{placeholder:"请输入手机号",clearable:""},model:{value:e.formData.phone,callback:function(t){e.$set(e.formData,"phone",t)},expression:"formData.phone"}})],1),o("el-form-item",{attrs:{label:"公卫账号",prop:"gwAccount"}},[o("el-input",{style:{width:"100%"},attrs:{placeholder:"请输入公卫账号",clearable:""},model:{value:e.formData.gwAccount,callback:function(t){e.$set(e.formData,"gwAccount",t)},expression:"formData.gwAccount"}})],1),o("el-form-item",{attrs:{label:"公卫密码",prop:"gwPassword"}},[o("el-input",{style:{width:"100%"},attrs:{placeholder:"请输入公卫密码",clearable:""},model:{value:e.formData.gwPassword,callback:function(t){e.$set(e.formData,"gwPassword",t)},expression:"formData.gwPassword"}})],1),o("el-form-item",{attrs:{label:"所属区域",prop:"areaCode"}},[o("el-select",{style:{width:"100%"},attrs:{placeholder:"请选择所属区域",clearable:"",filterable:""},model:{value:e.formData.areaCode,callback:function(t){e.$set(e.formData,"areaCode",t)},expression:"formData.areaCode"}},e._l(e.areaCodeOptions,(function(e,t){return o("el-option",{key:t,attrs:{label:e.Name+" ["+e.Code+"]",value:e.Code,disabled:e.disabled}})})),1)],1),o("el-form-item",{attrs:{size:"large"}},[o("el-button",{attrs:{type:"primary"},on:{click:e.submitForm}},[e._v("检测")]),o("el-button",{on:{click:e.resetForm}},[e._v("重置")])],1)],1),o("div",[o("div",{staticClass:"th-msg-title"},[e._v("错误信息:")]),e.error&&e.error.message?o("div",{staticClass:"th-msg-detail",staticStyle:{color:"red"}},[e._v(" "+e._s(e.error.message)+" "),o("br"),e.error.errorHandleButtonShow?o("el-button",{attrs:{type:"primary",size:"mini"},on:{click:e.resolveError}},[e._v(" "+e._s(e.error.errorHandleButtonText)+" ")]):e._e()],1):e._e(),o("div",{staticClass:"th-msg-title",staticStyle:{"margin-bottom":"7px"}},[e._v("医生信息:")]),e.result&&e.result.DoctorInfo?o("div",{staticClass:"th-msg-detail"},[e._v(" 医生名字: "+e._s(e.result.DoctorInfo.doctorName)+" "),o("br"),e._v(" 手机号: "+e._s(e.result.DoctorInfo.phoneNo)+" "),o("br"),o("span",{style:{color:50021===e.error.code||50022===e.error.code?"red":""}},[e._v(" 公卫账号: "+e._s(e.result.DoctorInfo.jobNumber)+" ")]),o("br"),o("span",{style:{color:50011===e.error.code?"red":""}},[e._v(" 卫生院: "+e._s(e.result.DoctorInfo.hospitalName)+" "+e._s(e.result.DoctorInfo.hospitalCode)+" ")]),o("br")]):e._e(),e.result&&e.result.DoctorInfo?o("div",[o("el-button",{attrs:{type:"primary",size:"mini"},on:{click:e.updateDoctorInfo}},[e._v(" 修改医生信息 ")]),o("el-button",{attrs:{type:"primary",size:"mini"},on:{click:e.resetYixinPassword}},[e._v(" 重置医信密码 ")]),o("el-button",{attrs:{type:"danger",size:"mini"},on:{click:e.deleteDoctor}},[e._v(" 删除医生 ")])],1):e._e(),o("div",{staticClass:"th-msg-title"},[e._v(" 角色信息 ")]),e.result&&e.result.DoctorLoginInfo?o("div",{staticClass:"th-msg-detail",style:{color:50031===e.error.code?"red":""}},[e._v(" "+e._s(e.result.DoctorLoginInfo.roleName)+" 【"+e._s(e.result.DoctorLoginInfo.roleCode)+"】 ")]):e._e()]),o("create-account-dialog",{ref:"createAccountDialogRef",attrs:{visible:e.dialogVisible.createAccount,title:"新建用户或完善信息","update-id":e.dialogVisible.createAccountProps.updateId,"update-field-limit":e.dialogVisible.createAccountProps.updateFieldLimit},on:{"update:visible":function(t){return e.$set(e.dialogVisible,"createAccount",t)},reLogin:e.reLogin,clearError:e.clearError}})],1)},n=[],s=(o("d3b7"),o("caad"),o("b5f9")),l=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",[o("el-dialog",e._g(e._b({on:{open:e.onOpen,close:e.onClose}},"el-dialog",e.$attrs,!1),e.$listeners),[o("el-form",{ref:"elForm",attrs:{model:e.formData,rules:e.rules,size:"small","label-width":"100px"}},[e.updateFieldLimit&&"userName"!==e.updateFieldLimit&&e.formDataOnInit.userName?e._e():o("el-form-item",{attrs:{label:"姓名",prop:"userName"}},[o("el-input",{style:{width:"100%"},attrs:{placeholder:"请输入姓名",clearable:""},model:{value:e.formData.userName,callback:function(t){e.$set(e.formData,"userName",t)},expression:"formData.userName"}})],1),e.updateFieldLimit&&"phone"!==e.updateFieldLimit&&e.formDataOnInit.phone?e._e():o("el-form-item",{attrs:{label:"手机号",prop:"phone"}},[o("el-input",{style:{width:"100%"},attrs:{placeholder:"请输入手机号",clearable:"",disabled:!!e.formDataOnInit.phone},model:{value:e.formData.phone,callback:function(t){e.$set(e.formData,"phone",t)},expression:"formData.phone"}})],1),e.updateFieldLimit&&"gwAccount"!==e.updateFieldLimit&&e.formDataOnInit.gwAccount?e._e():o("el-form-item",{attrs:{label:"公卫账号",prop:"gwAccount"}},[o("el-input",{style:{width:"100%"},attrs:{placeholder:"请输入公卫账号",clearable:""},model:{value:e.formData.gwAccount,callback:function(t){e.$set(e.formData,"gwAccount",t)},expression:"formData.gwAccount"}})],1),e.updateFieldLimit&&"hospitalCode"!==e.updateFieldLimit&&e.formDataOnInit.hospitalCode?e._e():o("el-form-item",{attrs:{label:"所属卫生院",prop:"hospitalCode"}},[o("el-select",{style:{width:"100%"},attrs:{placeholder:"请选择所属卫生院",clearable:"",filterable:""},model:{value:e.formData.hospitalCode,callback:function(t){e.$set(e.formData,"hospitalCode",t)},expression:"formData.hospitalCode"}},e._l(e.hospitalCodeOptions,(function(e,t){return o("el-option",{key:t,attrs:{label:e.Name+" ["+e.Code+"]",value:e.Code,disabled:e.disabled}})})),1)],1),e.updateFieldLimit&&"sex"!==e.updateFieldLimit?e._e():o("el-form-item",{attrs:{label:"性别",prop:"sex"}},[o("el-radio-group",{attrs:{size:"medium"},model:{value:e.formData.sex,callback:function(t){e.$set(e.formData,"sex",t)},expression:"formData.sex"}},e._l(e.sexOptions,(function(t,r){return o("el-radio",{key:r,attrs:{label:t.value,disabled:t.disabled}},[e._v(e._s(t.label)+" ")])})),1)],1)],1),o("div",{attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:e.close}},[e._v("取消")]),o("el-button",{attrs:{type:"primary"},on:{click:e.handelConfirm}},[e._v("确定")])],1)],1)],1)},c=[],u={name:"CreateAccountDialog",components:{},inheritAttrs:!1,props:{updateId:{type:String,default:""},updateFieldLimit:{type:String,default:null}},data:function(){return{hospitalCodeLoading:!1,formData:{userName:"",phone:"",gwAccount:void 0,hospitalCode:void 0,sex:1},formDataOnInit:{userName:"",phone:"",gwAccount:void 0,hospitalCode:void 0,sex:1},hospitalCodeOptions:[],sexOptions:[{label:"男",value:1},{label:"女",value:2}],doctorBaseDataForUpdate:{}}},computed:{rules:function(){return{userName:[{required:!0,message:"请输入姓名",trigger:"blur"}],phone:[{required:!0,message:"请输入手机号",trigger:"blur"}],gwAccount:[{required:!0,message:"请输入公卫账号",trigger:"blur"}],hospitalCode:[{required:!0,message:"请选择所属卫生室",trigger:"change"}],sex:[]}}},watch:{},created:function(){},mounted:function(){var e=this;Object(s["g"])("").then((function(t){console.log(t),e.hospitalCodeOptions=t.data}))},methods:{onOpen:function(){},onClose:function(){this.$refs["elForm"].resetFields()},close:function(){this.$emit("update:visible",!1)},handelConfirm:function(){var e=this;this.$refs["elForm"].validate((function(t){t&&(e.loading=!0,e.formData.id=e.updateId,Object(s["b"])(e.formData).then((function(t){e.$message({type:"success",message:t.msg||"新建成功"}),e.close(),e.$emit("clearError")})).catch((function(t){40001===t.code&&e.$emit("reLogin")})).finally((function(){e.loading=!1})))}))},refUpdateUserInfoOnCreate:function(e,t,o,r){this.formData.userName=null!==e&&void 0!==e?e:"",this.formData.phone=null!==t&&void 0!==t?t:"",this.formData.gwAccount=null!==o&&void 0!==o?o:"",this.formData.hospitalCode=null!==r&&void 0!==r?r:"",this.formDataOnInit=JSON.parse(JSON.stringify(this.formData))},refUpdateUserInfoOnUpdate:function(e){var t,o,r,a,i=e.DoctorInfo;this.formData.userName=null!==(t=i.doctorName)&&void 0!==t?t:"",this.formData.phone=null!==(o=i.phoneNo)&&void 0!==o?o:"",this.formData.gwAccount=null!==(r=i.jobNumber)&&void 0!==r?r:"",this.formData.hospitalCode=null!==(a=i.hospitalCode)&&void 0!==a?a:"",this.formDataOnInit=JSON.parse(JSON.stringify(this.formData))}}},d=u,f=o("2877"),m=Object(f["a"])(d,l,c,!1,null,null,null),h=m.exports,p={name:"YixinAnalz",components:{CreateAccountDialog:h},props:[],data:function(){return{loading:!1,dialogVisible:{createAccount:!1,createAccountProps:{updateId:"",updateFieldLimit:""}},result:{},formData:{phone:"",gwAccount:void 0,gwPassword:void 0,areaCode:void 0},areaCodeOptions:[{Name:"山东青岛平度",Code:"370283"},{Name:"贵州安顺镇宁",Code:"520423"}],error:{message:"",code:0,errorHandleButtonShow:!1,errorHandleButtonText:"错误处理"}}},computed:{rules:function(){var e=!!this.formData.gwPassword||!!this.formData.areaCode;return{phone:[{required:!0,message:"请输入手机号",trigger:"blur"}],gwAccount:[{required:e,message:"请输入公卫账号",trigger:"blur"}],gwPassword:[{required:e,message:"请输入公卫密码",trigger:"blur"}],areaCode:[{required:e,message:"请选择区域",trigger:"change"}]}}},watch:{},created:function(){},mounted:function(){var e=this;this.loading=!0,Object(s["e"])().then((function(t){e.areaCodeOptions=t.data||[]})).finally((function(){e.loading=!1}))},methods:{reLogin:function(){this.$emit("reLogin")},clearError:function(){this.error={message:"",code:0,errorHandleButtonShow:!1,errorHandleButtonText:"错误处理"}},submitForm:function(){var e=this;11===this.formData.phone.length?(this.result={},this.clearError(),this.$refs["elForm"].validate((function(t){t&&(e.loading=!0,Object(s["a"])(e.formData).then((function(t){console.log(t),e.$message({message:"无异常",type:"success"}),e.result=t.data})).catch((function(t){console.log(t),e.handleErrData(t),t.data&&(e.result=t.data)})).finally((function(){e.loading=!1})))}))):this.$message({message:"填写的手机号需要11位",type:"error"})},handleErrData:function(e){40001!==e.code?(this.error.code=e.code,this.error.message=e.msg||"",50001===e.code&&(this.error.errorHandleButtonText="一键注册",this.error.errorHandleButtonShow=!0),50002===e.code&&(this.error.errorHandleButtonText="一键完善",this.error.errorHandleButtonShow=!0),50011===e.code&&(this.error.errorHandleButtonText="一键修改",this.error.errorHandleButtonShow=!0),50021===e.code&&(this.error.errorHandleButtonText="一键维护",this.error.errorHandleButtonShow=!0),50022===e.code&&(this.error.errorHandleButtonText="更新公卫账号",this.error.errorHandleButtonShow=!0),50031===e.code&&(this.error.errorHandleButtonText="修改为医生角色",this.error.errorHandleButtonShow=!0)):this.$emit("reLogin")},resetForm:function(){this.$refs["elForm"].resetFields()},resolveError:function(){var e=this;if(50001!==this.error.code&&50002!==this.error.code||(this.dialogVisible.createAccount=!0,this.dialogVisible.createAccountProps={isUpdate:!1,updateFieldLimit:""},this.$refs.createAccountDialogRef.refUpdateUserInfoOnCreate("",this.formData.phone,this.formData.gwAccount,"")),[50011,50021,50022].includes(this.error.code)){var t={50011:"hospitalCode",50021:"gwAccount",50022:"gwAccount"};if(this.dialogVisible.createAccount=!0,this.dialogVisible.createAccountProps={updateId:this.result.DoctorInfo.id+"",updateFieldLimit:t[this.error.code]},50022===this.error.code){var o=JSON.parse(JSON.stringify(this.result));o.DoctorInfo&&(o.DoctorInfo.jobNumber=this.formData.gwAccount),this.$refs.createAccountDialogRef.refUpdateUserInfoOnUpdate(o)}else this.$refs.createAccountDialogRef.refUpdateUserInfoOnUpdate(this.result)}50031===this.error.code&&this.$confirm("确认修改为医生角色？").then((function(t){e.loading=!0,Object(s["d"])(e.result.DoctorInfo.phoneNo).then((function(t){e.$message({message:t.msg,type:"success"}),e.submitForm()})).catch((function(e){})).finally((function(){e.loading=!1}))})).catch((function(e){}))},updateDoctorInfo:function(){this.dialogVisible.createAccount=!0,this.dialogVisible.createAccountProps={updateId:this.result.DoctorInfo.id+"",updateFieldLimit:""},this.$refs.createAccountDialogRef.refUpdateUserInfoOnUpdate(this.result)},deleteDoctor:function(){var e=this;this.result&&this.result.DoctorInfo&&this.$confirm("确认删除？").then((function(t){e.loading=!0,Object(s["c"])(e.result.DoctorInfo.id,e.result.DoctorInfo.phoneNo).then((function(t){e.$message({message:t.msg,type:"success"}),e.submitForm()})).catch((function(e){})).finally((function(){e.loading=!1}))})).catch((function(e){}))},resetYixinPassword:function(){var e=this;this.result&&this.result.DoctorInfo&&this.$confirm("确认重置密码为123456？").then((function(t){e.loading=!0,Object(s["i"])(e.result.DoctorInfo.id).then((function(t){e.$message({message:t.msg,type:"success"})})).catch((function(e){})).finally((function(){e.loading=!1}))})).catch((function(e){}))}}},g=p,b=(o("c971"),Object(f["a"])(g,i,n,!1,null,"2163b086",null)),v=b.exports,D=o("911b"),w={name:"Index",components:{Login:D["a"],YixinAnalz:v},data:function(){return{loginSuccess:!0}},methods:{reLogin:function(){var e=this;this.loginSuccess=!1,setTimeout((function(){e.$refs.loginFormRef.refReloadCodeImage()}),200)},handleClose:function(){this.loginSuccess=!1},loginSubmit:function(){this.$refs.loginFormRef.login()}}},C=w,y=(o("750e"),Object(f["a"])(C,r,a,!1,null,"799b8042",null));t["default"]=y.exports}}]);