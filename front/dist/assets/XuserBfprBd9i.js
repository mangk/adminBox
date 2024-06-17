import{I as h,r as m,a as F,J as le,f as r,o as C,c as T,d as l,w as o,g as _,j as N,k as te,e as ae,F as oe,y as ne,i as re}from"./XindexrDa97Uv6.js";import ie from"./XpremissionSetMDhdITY5.js";import{e as se}from"./XdepartmentQwUOs5On.js";import{e as de}from"./XroleB_-QFiVh.js";import"./X_plugin-vue_export-helperDlAUqK2U.js";const q=window.adminX?window.adminX:{},V=q.BackendPrefix?"/"+q.BackendPrefix:"",ue=(i=1,c=20,k={})=>h({url:V+"/sys/setting/user/page",method:"POST",data:{page:i,page_size:c,query:k}}),me=i=>h({url:V+"/sys/setting/user/getById",method:"POST",data:{id:i}}),pe=i=>h({url:V+"/sys/setting/user",method:"POST",data:i}),_e=i=>h({url:V+"/sys/setting/user",method:"PUT",data:i}),ce=i=>h({url:V+"/sys/setting/user",method:"DELETE",data:{id:i}}),fe={class:"main-content"},ge={class:"dialog-footer"},Ve={__name:"user",setup(i){const{proxy:c}=re(),k=m([]),f=m(1),g=m(20),x=m(0),P=m(0),v=m(!1),b=m(),a=F({id:0,username:"",phone:"",email:"",nick_name:"",avatar:"",department_ids:[],role_ids:[],enable:!0}),A=F({name:[{required:!0,message:"API名称 不能为空",trigger:"blur"}],path:[{required:!0,message:"API地址 不能为空",trigger:"blur"}],method:[{required:!0,message:"API方法 不能为空",trigger:"blur"}]}),I=m([]),$=m([]),w=()=>{ue(f.value,g.value).then(n=>{k.value=n.data.list,f.value=n.data.page,g.value=n.data.page_size,x.value=n.data.total})},E=n=>{g.value=n,w()},O=n=>{f.value=n,w()},D=(n=!1)=>{se().then(e=>{I.value=e.data}),de().then(e=>{$.value=e.data}),n?(a.id=n,me(n).then(e=>{e.code===0&&(a.id=e.data.id,a.username=e.data.username,a.phone=e.data.phone,a.email=e.data.email,a.nick_name=e.data.nick_name,a.avatar=e.data.avatar,a.enable=e.data.enable,a.department_ids=e.data.department_ids,a.role_ids=e.data.role_ids)})):a.id=0,v.value=!0},j=n=>{n?ce(n).then(e=>{e.code===0&&(c.$message.success(e.msg),w())}):c.$message.error("请选择数据ID"),console.log("del",n)},L=async n=>{n&&await n.validate((e,z)=>{if(!e)console.log("error submit!",z);else{let p={};console.log(a.id),a.id?p=_e(a):p=pe(a),p.then(s=>{s.code===0&&(c.$message.success(s.msg),w(),v.value=!1,U(n))})}})},S=n=>{v.value=!1,U(n)},U=n=>{n&&n.resetFields()};return le(()=>{w()}),(n,e)=>{const z=r("Plus"),p=r("el-icon"),s=r("el-button"),d=r("el-form-item"),u=r("el-table-column"),R=r("Edit"),X=r("Filter"),J=r("Delete"),M=r("el-popconfirm"),G=r("el-table"),H=r("el-pagination"),y=r("el-input"),K=r("el-cascader"),Q=r("el-option"),W=r("el-select"),Y=r("el-switch"),Z=r("el-form"),ee=r("el-dialog");return C(),T("div",fe,[l(d,{label:""},{default:o(()=>[l(s,{type:"primary",onClick:e[0]||(e[0]=t=>D())},{default:o(()=>[l(p,null,{default:o(()=>[l(z)]),_:1}),_(" 新增用户 ")]),_:1})]),_:1}),l(G,{data:k.value,"row-key":"id",height:"var(--global-table)","header-cell-class-name":"global-el-table-header",border:"","highlight-current-row":"","show-overflow-tooltip":""},{default:o(()=>[l(u,{prop:"id",label:"ID",sortable:"",fixed:""}),l(u,{prop:"uuid",label:"UUID",width:"300"}),l(u,{prop:"nick_name",label:"昵称",width:"120"}),l(u,{prop:"username",label:"用户名"}),l(u,{prop:"phone",label:"手机号",width:"160"}),l(u,{prop:"email",label:"Email",width:"220"}),l(u,{prop:"avatar",label:"头像",width:"120"}),l(u,{prop:"enable",label:"是否启用",width:"120"}),l(u,{fixed:"right",label:"操作",width:"200"},{default:o(t=>[l(s,{link:"",type:"primary",size:"small",onClick:B=>D(t.row.id)},{default:o(()=>[l(p,null,{default:o(()=>[l(R)]),_:1}),_(" 编辑 ")]),_:2},1032,["onClick"]),l(s,{link:"",type:"primary",size:"small",onClick:B=>P.value=t.row.id},{default:o(()=>[l(p,null,{default:o(()=>[l(X)]),_:1}),_(" 权限设置 ")]),_:2},1032,["onClick"]),t.row.children?te("",!0):(C(),N(M,{key:0,title:"删除后不可恢复，确定删除API【"+t.row.nickName+"("+t.row.userName+")】?",onConfirm:B=>j(t.row.id),width:"200"},{reference:o(()=>[l(s,{link:"",type:"primary",size:"small"},{default:o(()=>[l(p,null,{default:o(()=>[l(J)]),_:1}),_(" 删除 ")]),_:1})]),_:2},1032,["title","onConfirm"]))]),_:1})]),_:1},8,["data"]),l(H,{"current-page":f.value,"onUpdate:currentPage":e[1]||(e[1]=t=>f.value=t),"page-size":g.value,"onUpdate:pageSize":e[2]||(e[2]=t=>g.value=t),"page-sizes":[20,50,100,200],size:"small",layout:"total, sizes, prev, pager, next, jumper",total:x.value,onSizeChange:E,onCurrentChange:O},null,8,["current-page","page-size","total"]),l(ee,{modelValue:v.value,"onUpdate:modelValue":e[14]||(e[14]=t=>v.value=t),title:a.id?"编辑用户":"新建用户",width:"800",onClose:e[15]||(e[15]=t=>S(b.value))},{footer:o(()=>[ae("div",ge,[l(s,{onClick:e[11]||(e[11]=t=>U(b.value))},{default:o(()=>[_("重置")]),_:1}),l(s,{onClick:e[12]||(e[12]=t=>S(b.value))},{default:o(()=>[_("取消")]),_:1}),l(s,{type:"primary",onClick:e[13]||(e[13]=t=>L(b.value))},{default:o(()=>[_("保存")]),_:1})])]),default:o(()=>[l(Z,{ref_key:"formRef",ref:b,model:a,rules:A,"status-icon":"",inline:"","label-position":"top","show-all-levels":!1},{default:o(()=>[l(d,{label:"用户名",prop:"username",style:{width:"40%"}},{default:o(()=>[l(y,{modelValue:a.username,"onUpdate:modelValue":e[3]||(e[3]=t=>a.username=t)},null,8,["modelValue"])]),_:1}),l(d,{label:"手机号",prop:"phone",style:{width:"40%"}},{default:o(()=>[l(y,{modelValue:a.phone,"onUpdate:modelValue":e[4]||(e[4]=t=>a.phone=t)},null,8,["modelValue"])]),_:1}),l(d,{label:"Email",prop:"email",style:{width:"40%"}},{default:o(()=>[l(y,{modelValue:a.email,"onUpdate:modelValue":e[5]||(e[5]=t=>a.email=t)},null,8,["modelValue"])]),_:1}),l(d,{label:"昵称",prop:"nick_name",style:{width:"40%"}},{default:o(()=>[l(y,{modelValue:a.nick_name,"onUpdate:modelValue":e[6]||(e[6]=t=>a.nick_name=t)},null,8,["modelValue"])]),_:1}),l(d,{label:"所属部门",prop:"department_ids",style:{width:"40%"}},{default:o(()=>[l(K,{modelValue:a.department_ids,"onUpdate:modelValue":e[7]||(e[7]=t=>a.department_ids=t),options:I.value,placeholder:"设置用户部门",props:{checkStrictly:!0,expandTrigger:"hover",value:"id",label:"name",multiple:!0,emitPath:!1},style:{width:"100%"}},null,8,["modelValue","options"])]),_:1}),l(d,{label:"用户角色",prop:"role_ids",style:{width:"40%"}},{default:o(()=>[l(W,{modelValue:a.role_ids,"onUpdate:modelValue":e[8]||(e[8]=t=>a.role_ids=t),multiple:"",placeholder:"设置用户角色"},{default:o(()=>[(C(!0),T(oe,null,ne($.value,t=>(C(),N(Q,{key:t.id,label:t.name,value:t.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),l(d,{label:"头像",prop:"avatar",style:{width:"40%"}},{default:o(()=>[l(y,{modelValue:a.avatar,"onUpdate:modelValue":e[9]||(e[9]=t=>a.avatar=t)},null,8,["modelValue"])]),_:1}),l(d,{label:"是否启用",prop:"enable",style:{width:"40%"}},{default:o(()=>[l(Y,{modelValue:a.enable,"onUpdate:modelValue":e[10]||(e[10]=t=>a.enable=t),"inline-prompt":"","active-icon":"Check","inactive-icon":"Close"},null,8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])]),_:1},8,["modelValue","title"]),l(ie,{modelValue:P.value,"onUpdate:modelValue":e[16]||(e[16]=t=>P.value=t),type:"sys_user"},null,8,["modelValue"])])}}};export{Ve as default};
