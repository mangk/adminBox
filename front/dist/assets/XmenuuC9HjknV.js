import"./XvueCyDLjQZ5.js";import{a as e,b as l,c as a,d as t,m as o}from"./XmenuCOh4UbVh.js";import{Q as d}from"./X@element-plusCAKmCgXq.js";import{_ as i}from"./X_plugin-vue_export-helperBCo6x5W8.js";import{r,$ as u,j as s,ah as m,o as n,c as p,X as c,Q as f,V as h,a as _,W as b,P as v,S as g,U as w,F as y,a8 as V,K as j,i as k}from"./X@vueSdcp2x8a.js";import"./XindexCDFV15DS.js";import"./XpiniaDkaaaSbL.js";import"./Xelement-plusJnkMASPo.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCSz_Rkca.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./Xvue3-sfc-loaderDa78GGQm.js";import"./XaxiosCURSphCx.js";import"./XmittCNZ6avp8.js";import"./Xvue-routerCZMibURm.js";import"./XnprogressDgs0sPf-.js";const X={class:"main-content"},x={key:0,class:"icon-column"},U={key:0,style:{position:"absolute","z-index":"9999",padding:"4px 10px 0"}},C={style:{"text-align":"left",margin:"0 0 0 10px"}},q={class:"dialog-footer"},z=i({__name:"menu",setup(i){const{proxy:z}=k(),E=r(!1),$=r([]);r("");const D=u([]);for(const[e]of Object.entries(d))D.push({key:e,label:e});const F=r(),I=u({id:0,pid:"",name:"",path:"",hidden:!1,component:"",sort:0,meta:{title:"",keep_alive:!1,default_menu:!1,icon:"",auto_close:!1,sc_path:"",action_list:null}}),K=u({pid:[{required:!0,trigger:"blur",validator:(e,l,a)=>{""!==l?l!==I.id?a():a(new Error("所选菜单不能为当前编辑菜单")):a(new Error("父级菜单 不能为空"))}}],name:[{required:!0,message:"路由名称 不能为空",trigger:"blur"}],path:[{required:!0,message:"路由地址 不能为空",trigger:"blur"}],hidden:[{required:!0,message:"请选择 是否隐藏",trigger:"blur"}],"meta.title":[{required:!0,message:"菜单名称 不能为空",trigger:"blur"}],"meta.keep_alive":[{required:!0,message:"请选择 是否KeepAlive",trigger:"blur"}],"meta.default_menu":[{required:!0,message:"请选择 是否为基础页面",trigger:"blur"}],"meta.auto_close":[{required:!0,message:"请选择 是否自动关闭标签",trigger:"blur"}]}),P=e=>{E.value=!1,A(e)},A=e=>{e&&e.resetFields()},Q=(e=!1)=>{e?(I.id=e,a(e).then((e=>{0===e.code&&(I.pid=e.data.pid,I.meta.title=e.data.meta.title,I.name=e.data.name,I.path=e.data.path,I.hidden=e.data.hidden,I.component=e.data.component,I.sort=e.data.sort,I.meta=e.data.meta)}))):I.id=0,M(),E.value=!0},S=()=>{o().then((e=>{$.value=e.data[0].children,G.value=e.data}))},G=r([]),M=()=>{O(G.value,!1)},O=(e,l)=>{e&&e.forEach((e=>{e.children&&e.children.length?(e.title=e.meta.title,e.disabled=(l||e.id===I.id)&&I.id,O(e.children,e.disabled)):(e.title=e.meta.title,e.disabled=(l||e.id===I.id)&&I.id)}))};return s((()=>{S()})),(a,o)=>{const d=m("Plus"),i=m("el-icon"),r=m("el-button"),u=m("el-form-item"),s=m("el-table-column"),k=m("Edit"),M=m("Cherry"),O=m("Delete"),R=m("el-popconfirm"),T=m("el-table"),W=m("el-cascader"),B=m("el-input"),H=m("el-option"),J=m("el-select"),L=m("el-form"),N=m("el-dialog");return n(),p("div",X,[c(u,{label:""},{default:f((()=>[c(r,{type:"primary",onClick:o[0]||(o[0]=e=>Q())},{default:f((()=>[c(i,null,{default:f((()=>[c(d)])),_:1}),h(" 新建目录 ")])),_:1})])),_:1}),c(T,{data:$.value,"row-key":"id",height:"var(--global-table)",border:"","highlight-current-row":"","show-overflow-tooltip":"","default-expand-all":""},{default:f((()=>[c(s,{prop:"id",label:"ID",sortable:"",fixed:""}),c(s,{label:"菜单名称","min-width":"120"},{default:f((e=>[_("span",null,b(e.row.meta.title),1)])),_:1}),c(s,{label:"图标","min-width":"140"},{default:f((e=>[e.row.meta.icon?(n(),p("div",x,[c(i,null,{default:f((()=>[(n(),v(g(e.row.meta.icon)))])),_:2},1024),_("span",null,b(e.row.meta.icon),1)])):w("",!0)])),_:1}),c(s,{prop:"name",label:"路由名称",width:"120"}),c(s,{prop:"path",label:"路由地址",width:"120"}),c(s,{prop:"hidden",label:"隐藏"}),c(s,{prop:"pid",label:"父节点"}),c(s,{prop:"sort",label:"排序",width:"120",sortable:""}),c(s,{prop:"component",label:"模版路经",width:"260"}),c(s,{label:"服务端模版",width:"260"},{default:f((e=>[h(b(e.row.meta?e.row.meta.sc_path:""),1)])),_:1}),c(s,{prop:"action_list",label:"动作列表",width:"260"}),c(s,{fixed:"right",label:"操作",width:"210"},{default:f((e=>[c(r,{link:"",type:"primary",size:"small",onClick:l=>Q(e.row.id)},{default:f((()=>[c(i,null,{default:f((()=>[c(k)])),_:1}),h(" 编辑 ")])),_:2},1032,["onClick"]),c(r,{link:"",type:"primary",size:"small"},{default:f((()=>[c(i,null,{default:f((()=>[c(M)])),_:1}),h(" 编辑动作 ")])),_:1}),e.row.children?w("",!0):(n(),v(R,{key:0,title:"删除后不可恢复，确定删除菜单【"+e.row.meta.title+"】?",onConfirm:l=>{var a;(a=e.row.id)?t(a).then((e=>{0===e.code&&(z.$message.success(e.msg),S())})):z.$message.error("请选择数据ID")},width:"200"},{reference:f((()=>[c(r,{link:"",type:"primary",size:"small"},{default:f((()=>[c(i,null,{default:f((()=>[c(O)])),_:1}),h(" 删除 ")])),_:1})])),_:2},1032,["title","onConfirm"]))])),_:1})])),_:1},8,["data"]),c(N,{modelValue:E.value,"onUpdate:modelValue":o[16]||(o[16]=e=>E.value=e),title:I.id?"编辑菜单":"新建菜单",width:"800",onClose:o[17]||(o[17]=e=>P(F.value))},{footer:f((()=>[_("div",q,[c(r,{onClick:o[13]||(o[13]=e=>A(F.value))},{default:f((()=>[h("重置")])),_:1}),c(r,{onClick:o[14]||(o[14]=e=>P(F.value))},{default:f((()=>[h("取消")])),_:1}),c(r,{type:"primary",onClick:o[15]||(o[15]=a=>(async a=>{a&&await a.validate(((t,o)=>{if(t){let t={};t=I.id?e(I):l(I),t.then((e=>{0===e.code&&(z.$message.success(e.msg),S(),E.value=!1,A(a))}))}}))})(F.value))},{default:f((()=>[h("保存")])),_:1})])])),default:f((()=>[c(L,{ref_key:"formRef",ref:F,model:I,rules:K,"status-icon":"","label-position":"top","show-all-levels":!1,inline:""},{default:f((()=>[c(u,{label:"父级菜单",prop:"pid",style:{width:"40%"}},{default:f((()=>[c(W,{modelValue:I.pid,"onUpdate:modelValue":o[1]||(o[1]=e=>I.pid=e),options:G.value,props:{expandTrigger:"hover",checkStrictly:!0,value:"id",label:"title",emitPath:!1},style:{width:"100%"}},null,8,["modelValue","options"])])),_:1}),c(u,{label:"菜单名称",prop:"meta.title",style:{width:"40%"}},{default:f((()=>[c(B,{modelValue:I.meta.title,"onUpdate:modelValue":o[2]||(o[2]=e=>I.meta.title=e)},null,8,["modelValue"])])),_:1}),c(u,{label:"路由名称(name)",prop:"name",style:{width:"40%"}},{default:f((()=>[c(B,{modelValue:I.name,"onUpdate:modelValue":o[3]||(o[3]=e=>I.name=e)},null,8,["modelValue"])])),_:1}),c(u,{label:"路由地址(path)",prop:"path",style:{width:"40%"}},{default:f((()=>[c(B,{modelValue:I.path,"onUpdate:modelValue":o[4]||(o[4]=e=>I.path=e)},null,8,["modelValue"])])),_:1}),c(u,{label:"图标",prop:"meta.icon",style:{width:"40%"}},{default:f((()=>[I.meta.icon?(n(),p("span",U,[c(i,null,{default:f((()=>[(n(),v(g(I.meta.icon)))])),_:1})])):w("",!0),c(J,{modelValue:I.meta.icon,"onUpdate:modelValue":o[5]||(o[5]=e=>I.meta.icon=e),clearable:"",class:"icon-select"},{default:f((()=>[(n(!0),p(y,null,V(D,(e=>(n(),v(H,{key:e.key,label:e.key,value:e.label},{default:f((()=>[_("span",{class:j(["gva-icon",e.label]),style:{padding:"3px 0 0"}},[c(i,null,{default:f((()=>[(n(),v(g(e.label)))])),_:2},1024)],2),_("span",C,b(e.key),1)])),_:2},1032,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),c(u,{label:"排序",prop:"sort",style:{width:"40%"}},{default:f((()=>[c(B,{type:"number",modelValue:I.sort,"onUpdate:modelValue":o[6]||(o[6]=e=>I.sort=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),c(u,{label:"模版地址",prop:"meta.component",style:{width:"40%"}},{default:f((()=>[c(B,{modelValue:I.component,"onUpdate:modelValue":o[7]||(o[7]=e=>I.component=e),placeholder:"views/util/serverComponent.vue"},null,8,["modelValue"])])),_:1}),c(u,{label:"服务端模版地址",prop:"meta.sc_path",style:{width:"40%"}},{default:f((()=>[c(B,{modelValue:I.meta.sc_path,"onUpdate:modelValue":o[8]||(o[8]=e=>I.meta.sc_path=e)},null,8,["modelValue"])])),_:1}),c(u,{label:"是否隐藏",prop:"hidden",style:{width:"40%"}},{default:f((()=>[c(J,{modelValue:I.hidden,"onUpdate:modelValue":o[9]||(o[9]=e=>I.hidden=e)},{default:f((()=>[c(H,{label:"否",value:!1}),c(H,{label:"是",value:!0})])),_:1},8,["modelValue"])])),_:1}),c(u,{label:"KeepAlive",prop:"meta.keep_alive",style:{width:"40%"}},{default:f((()=>[c(J,{modelValue:I.meta.keep_alive,"onUpdate:modelValue":o[10]||(o[10]=e=>I.meta.keep_alive=e)},{default:f((()=>[c(H,{label:"否",value:!1}),c(H,{label:"是",value:!0})])),_:1},8,["modelValue"])])),_:1}),c(u,{label:"基础页面",prop:"meta.default_menu",style:{width:"40%"}},{default:f((()=>[c(J,{modelValue:I.meta.default_menu,"onUpdate:modelValue":o[11]||(o[11]=e=>I.meta.default_menu=e)},{default:f((()=>[c(H,{label:"否",value:!1}),c(H,{label:"是",value:!0})])),_:1},8,["modelValue"])])),_:1}),c(u,{label:"自动关闭标签",prop:"meta.auto_close",style:{width:"40%"}},{default:f((()=>[c(J,{modelValue:I.meta.auto_close,"onUpdate:modelValue":o[12]||(o[12]=e=>I.meta.auto_close=e)},{default:f((()=>[c(H,{label:"否",value:!1}),c(H,{label:"是",value:!0})])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"])])}}},[["__scopeId","data-v-af8bbf3c"]]);export{z as default};
