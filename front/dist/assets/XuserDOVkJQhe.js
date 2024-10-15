import"./XvueBaz_cZdj.js";import{h as e}from"./XindexDwx5IuIZ.js";import l from"./XpremissionSetDLs14YYt.js";import{e as a}from"./XdepartmentCMP1Khh9.js";import{e as t}from"./XroleL3ucvj_U.js";import o from"./XfileUpload13MGz1lP.js";import{b as r,$ as i,j as d,ah as s,o as u,c as n,X as p,Q as m,V as h,P as c,U as v,a as f,F as _,a8 as g,i as b}from"./X@vuen9TiMG2O.js";import"./XpiniaYeQOi8UL.js";import"./Xelement-plusCLM398fa.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCCX6pF6_.js";import"./X@element-plusBnJF7_W-.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./XaxiosByKc3Ecy.js";import"./XmittCNZ6avp8.js";import"./Xvue-routerzsXVQF6n.js";import"./XnprogressDgs0sPf-.js";import"./Xvue3-sfc-loaderDa78GGQm.js";import"./X_plugin-vue_export-helperBCo6x5W8.js";const y={class:"main-content"},w={class:"dialog-footer"},V={__name:"user",setup(V){const{proxy:j}=b(),k=r([]),X=r(1),x=r(20),U=r(0),C=r(0),z=r(!1),P=r(),I=i({id:0,username:"",phone:"",email:"",nick_name:"",avatar:"",department_ids:[],role_ids:[],enable:!0}),E=i({name:[{required:!0,message:"API名称 不能为空",trigger:"blur"}],path:[{required:!0,message:"API地址 不能为空",trigger:"blur"}],method:[{required:!0,message:"API方法 不能为空",trigger:"blur"}]}),S=r([]),D=r([]),T=()=>{((l=1,a=20,t={})=>e({url:"/sys/setting/user/page",method:"POST",data:{page:l,page_size:a,query:t}}))(X.value,x.value).then((e=>{k.value=e.data.list,X.value=e.data.page,x.value=e.data.page_size,U.value=e.data.total}))},q=e=>{x.value=e,T()},A=e=>{X.value=e,T()},$=(l=!1)=>{a().then((e=>{S.value=e.data})),t().then((e=>{D.value=e.data})),l?(I.id=l,(l=>e({url:"/sys/setting/user/getById",method:"POST",data:{id:l}}))(l).then((e=>{0===e.code&&(I.id=e.data.id,I.username=e.data.username,I.phone=e.data.phone,I.email=e.data.email,I.nick_name=e.data.nick_name,I.avatar=e.data.avatar,I.enable=e.data.enable,I.department_ids=e.data.department_ids,I.role_ids=e.data.role_ids)}))):I.id=0,z.value=!0},F=l=>{l?(l=>e({url:"/sys/setting/user",method:"DELETE",data:{id:l}}))(l).then((e=>{0===e.code&&(j.$message.success(e.msg),T())})):j.$message.error("请选择数据ID")},O=async l=>{l&&await l.validate(((a,t)=>{if(a){let a={};a=I.id?e({url:"/sys/setting/user",method:"PUT",data:I}):(l=>e({url:"/sys/setting/user",method:"POST",data:l}))(I),a.then((e=>{0===e.code&&(j.$message.success(e.msg),T(),z.value=!1,N(l))}))}else;}))},B=e=>{z.value=!1,N(e)},N=e=>{e&&e.resetFields()};return d((()=>{T()})),(e,a)=>{const t=s("Plus"),r=s("el-icon"),i=s("el-button"),d=s("el-form-item"),b=s("el-table-column"),V=s("Edit"),j=s("Filter"),T=s("Delete"),L=s("el-popconfirm"),Q=s("el-table"),R=s("el-pagination"),G=s("el-input"),H=s("el-cascader"),J=s("el-option"),K=s("el-select"),M=s("el-image"),W=s("el-switch"),Y=s("el-form"),Z=s("el-dialog");return u(),n("div",y,[p(d,{label:""},{default:m((()=>[p(i,{type:"primary",onClick:a[0]||(a[0]=e=>$())},{default:m((()=>[p(r,null,{default:m((()=>[p(t)])),_:1}),h(" 新增用户 ")])),_:1})])),_:1}),p(Q,{data:k.value,"row-key":"id",height:"var(--global-table)",border:"","highlight-current-row":"","show-overflow-tooltip":""},{default:m((()=>[p(b,{prop:"id",label:"ID",sortable:"",fixed:""}),p(b,{prop:"uuid",label:"UUID",width:"300"}),p(b,{prop:"nick_name",label:"昵称",width:"120"}),p(b,{prop:"username",label:"用户名"}),p(b,{prop:"phone",label:"手机号",width:"160"}),p(b,{prop:"email",label:"Email",width:"220"}),p(b,{prop:"avatar",label:"头像",width:"120"}),p(b,{prop:"enable",label:"是否启用",width:"120"}),p(b,{fixed:"right",label:"操作",width:"200"},{default:m((e=>[p(i,{link:"",type:"primary",size:"small",onClick:l=>$(e.row.id)},{default:m((()=>[p(r,null,{default:m((()=>[p(V)])),_:1}),h(" 编辑 ")])),_:2},1032,["onClick"]),p(i,{link:"",type:"primary",size:"small",onClick:l=>C.value=e.row.id},{default:m((()=>[p(r,null,{default:m((()=>[p(j)])),_:1}),h(" 权限设置 ")])),_:2},1032,["onClick"]),e.row.children?v("",!0):(u(),c(L,{key:0,title:"删除后不可恢复，确定删除API【"+e.row.nickName+"("+e.row.userName+")】?",onConfirm:l=>F(e.row.id),width:"200"},{reference:m((()=>[p(i,{link:"",type:"primary",size:"small"},{default:m((()=>[p(r,null,{default:m((()=>[p(T)])),_:1}),h(" 删除 ")])),_:1})])),_:2},1032,["title","onConfirm"]))])),_:1})])),_:1},8,["data"]),p(R,{"current-page":X.value,"onUpdate:currentPage":a[1]||(a[1]=e=>X.value=e),"page-size":x.value,"onUpdate:pageSize":a[2]||(a[2]=e=>x.value=e),"page-sizes":[20,50,100,200],size:"small",layout:"total, sizes, prev, pager, next, jumper",total:U.value,onSizeChange:q,onCurrentChange:A},null,8,["current-page","page-size","total"]),p(Z,{modelValue:z.value,"onUpdate:modelValue":a[14]||(a[14]=e=>z.value=e),title:I.id?"编辑用户":"新建用户",width:"80vw","append-to-body":"",onClose:a[15]||(a[15]=e=>B(P.value))},{footer:m((()=>[f("div",w,[p(i,{onClick:a[11]||(a[11]=e=>N(P.value))},{default:m((()=>[h("重置")])),_:1}),p(i,{onClick:a[12]||(a[12]=e=>B(P.value))},{default:m((()=>[h("取消")])),_:1}),p(i,{type:"primary",onClick:a[13]||(a[13]=e=>O(P.value))},{default:m((()=>[h("保存")])),_:1})])])),default:m((()=>[p(Y,{ref_key:"formRef",ref:P,model:I,rules:E,"status-icon":"",inline:"","label-position":"top","show-all-levels":!1},{default:m((()=>[p(d,{label:"用户名",prop:"username",style:{width:"40%"}},{default:m((()=>[p(G,{modelValue:I.username,"onUpdate:modelValue":a[3]||(a[3]=e=>I.username=e)},null,8,["modelValue"])])),_:1}),p(d,{label:"手机号",prop:"phone",style:{width:"40%"}},{default:m((()=>[p(G,{modelValue:I.phone,"onUpdate:modelValue":a[4]||(a[4]=e=>I.phone=e)},null,8,["modelValue"])])),_:1}),p(d,{label:"Email",prop:"email",style:{width:"40%"}},{default:m((()=>[p(G,{modelValue:I.email,"onUpdate:modelValue":a[5]||(a[5]=e=>I.email=e)},null,8,["modelValue"])])),_:1}),p(d,{label:"昵称",prop:"nick_name",style:{width:"40%"}},{default:m((()=>[p(G,{modelValue:I.nick_name,"onUpdate:modelValue":a[6]||(a[6]=e=>I.nick_name=e)},null,8,["modelValue"])])),_:1}),p(d,{label:"所属部门",prop:"department_ids",style:{width:"40%"}},{default:m((()=>[p(H,{modelValue:I.department_ids,"onUpdate:modelValue":a[7]||(a[7]=e=>I.department_ids=e),options:S.value,placeholder:"设置用户部门",props:{checkStrictly:!0,expandTrigger:"hover",value:"id",label:"name",multiple:!0,emitPath:!1},style:{width:"100%"}},null,8,["modelValue","options"])])),_:1}),p(d,{label:"用户角色",prop:"role_ids",style:{width:"40%"}},{default:m((()=>[p(K,{modelValue:I.role_ids,"onUpdate:modelValue":a[8]||(a[8]=e=>I.role_ids=e),multiple:"",placeholder:"设置用户角色"},{default:m((()=>[(u(!0),n(_,null,g(D.value,(e=>(u(),c(J,{key:e.id,label:e.name,value:e.id},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),p(d,{label:"头像",prop:"avatar",style:{width:"40%"}},{default:m((()=>[p(M,{style:{width:"75px",height:"75px",border:"1px dashed var(--el-border-color)","border-radius":"5px"},src:I.avatar,fit:"contain"},null,8,["src"]),p(o,{modelValue:I.avatar,"onUpdate:modelValue":a[9]||(a[9]=e=>I.avatar=e),onlyPath:!0},{default:m((()=>[p(r,{style:{width:"75px",height:"75px",border:"1px dashed var(--el-border-color)","border-radius":"5px","margin-left":"10px"}},{default:m((()=>[p(t)])),_:1})])),_:1},8,["modelValue"])])),_:1}),p(d,{label:"是否启用",prop:"enable",style:{width:"40%"}},{default:m((()=>[p(W,{modelValue:I.enable,"onUpdate:modelValue":a[10]||(a[10]=e=>I.enable=e),"inline-prompt":"","active-icon":"Check","inactive-icon":"Close"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"]),p(l,{modelValue:C.value,"onUpdate:modelValue":a[16]||(a[16]=e=>C.value=e),type:"sys_user"},null,8,["modelValue"])])}}};export{V as default};