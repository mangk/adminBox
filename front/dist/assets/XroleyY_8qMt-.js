import"./XvueCyDLjQZ5.js";import{r as e,a as l,b as a,c as t,d as o}from"./XroleCzMtpx9h.js";import i from"./XpremissionSetBGGTYc5f.js";import{r,$ as s,j as u,ah as d,o as n,c as p,X as m,Q as c,V as f,P as v,U as g,a as j,i as _}from"./X@vueSdcp2x8a.js";import"./XindexCDFV15DS.js";import"./XpiniaDkaaaSbL.js";import"./Xelement-plusJnkMASPo.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCSz_Rkca.js";import"./X@element-plusCAKmCgXq.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./Xvue3-sfc-loaderDa78GGQm.js";import"./XaxiosCURSphCx.js";import"./XmittCNZ6avp8.js";import"./Xvue-routerCZMibURm.js";import"./XnprogressDgs0sPf-.js";import"./X_plugin-vue_export-helperBCo6x5W8.js";const h={class:"main-content"},X={class:"dialog-footer"},y={__name:"role",setup(y){const b=r(0),{proxy:w}=_(),k=r([]),z=r(1),C=r(20),V=r(0),x=r(!1),U=r(),I=s({id:0,name:"",description:""}),P=s({name:[{required:!0,message:"角色名称 不能为空",trigger:"blur"}]}),$=()=>{e(z.value,C.value).then((e=>{k.value=e.data.list,z.value=e.data.page,C.value=e.data.page_size,V.value=e.data.total}))},D=e=>{C.value=e,$()},F=e=>{z.value=e,$()},S=(e=!1)=>{e?(I.id=e,l(e).then((e=>{0===e.code&&(I.id=e.data.id,I.name=e.data.name,I.description=e.data.description)}))):I.id=0,x.value=!0},q=e=>{x.value=!1,A(e)},A=e=>{e&&e.resetFields()};return u((()=>{$()})),(e,l)=>{const r=d("Plus"),s=d("el-icon"),u=d("el-button"),_=d("el-form-item"),y=d("el-table-column"),E=d("Edit"),Q=d("Filter"),R=d("Delete"),B=d("el-popconfirm"),G=d("el-table"),H=d("el-pagination"),J=d("el-input"),K=d("el-form"),L=d("el-dialog");return n(),p("div",h,[m(_,{label:""},{default:c((()=>[m(u,{type:"primary",onClick:l[0]||(l[0]=e=>S())},{default:c((()=>[m(s,null,{default:c((()=>[m(r)])),_:1}),f(" 新建角色 ")])),_:1})])),_:1}),m(G,{data:k.value,"row-key":"id",height:"var(--global-table)",border:"","highlight-current-row":"","show-overflow-tooltip":""},{default:c((()=>[m(y,{prop:"id",label:"ID",sortable:"",fixed:"",width:"80"}),m(y,{prop:"name",label:"角色名称",width:"260"}),m(y,{prop:"description",label:"角色描述"}),m(y,{fixed:"right",label:"操作",width:"200"},{default:c((e=>[m(u,{link:"",type:"primary",size:"small",onClick:l=>S(e.row.id)},{default:c((()=>[m(s,null,{default:c((()=>[m(E)])),_:1}),f(" 编辑 ")])),_:2},1032,["onClick"]),m(u,{link:"",type:"primary",size:"small",onClick:l=>b.value=e.row.id},{default:c((()=>[m(s,null,{default:c((()=>[m(Q)])),_:1}),f(" 权限设置 ")])),_:2},1032,["onClick"]),e.row.children?g("",!0):(n(),v(B,{key:0,title:"删除后不可恢复，确定删除API【"+e.row.name+"】?",onConfirm:l=>{var t;(t=e.row.id)?a(t).then((e=>{0===e.code&&(w.$message.success(e.msg),$())})):w.$message.error("请选择数据ID")},width:"200"},{reference:c((()=>[m(u,{link:"",type:"primary",size:"small"},{default:c((()=>[m(s,null,{default:c((()=>[m(R)])),_:1}),f(" 删除 ")])),_:1})])),_:2},1032,["title","onConfirm"]))])),_:1})])),_:1},8,["data"]),m(H,{"current-page":z.value,"onUpdate:currentPage":l[1]||(l[1]=e=>z.value=e),"page-size":C.value,"onUpdate:pageSize":l[2]||(l[2]=e=>C.value=e),"page-sizes":[20,50,100,200],size:"small",layout:"total, sizes, prev, pager, next, jumper",total:V.value,onSizeChange:D,onCurrentChange:F},null,8,["current-page","page-size","total"]),m(L,{modelValue:x.value,"onUpdate:modelValue":l[8]||(l[8]=e=>x.value=e),title:I.id?"编辑角色":"新建角色",width:"800",onClose:l[9]||(l[9]=e=>q(U.value))},{footer:c((()=>[j("div",X,[m(u,{onClick:l[5]||(l[5]=e=>A(U.value))},{default:c((()=>[f("重置")])),_:1}),m(u,{onClick:l[6]||(l[6]=e=>q(U.value))},{default:c((()=>[f("取消")])),_:1}),m(u,{type:"primary",onClick:l[7]||(l[7]=e=>(async e=>{e&&await e.validate(((l,a)=>{if(l){let l={};l=I.id?t(I):o(I),l.then((l=>{0===l.code&&(w.$message.success(l.msg),$(),x.value=!1,A(e))}))}}))})(U.value))},{default:c((()=>[f("保存")])),_:1})])])),default:c((()=>[m(K,{ref_key:"formRef",ref:U,model:I,rules:P,"status-icon":"","label-position":"top","show-all-levels":!1},{default:c((()=>[m(_,{label:"角色名称",prop:"name"},{default:c((()=>[m(J,{modelValue:I.name,"onUpdate:modelValue":l[3]||(l[3]=e=>I.name=e)},null,8,["modelValue"])])),_:1}),m(_,{label:"角色描述",prop:"description"},{default:c((()=>[m(J,{modelValue:I.description,"onUpdate:modelValue":l[4]||(l[4]=e=>I.description=e)},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"]),m(i,{modelValue:b.value,"onUpdate:modelValue":l[10]||(l[10]=e=>b.value=e),type:"sys_role"},null,8,["modelValue"])])}}};export{y as default};
