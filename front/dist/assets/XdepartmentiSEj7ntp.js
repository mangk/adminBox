import{r as c,a as h,M as T,o as L,c as r,d as I,e as A,f as t,w as n,h as u,k as G,l as H,g as J,j as K}from"./XindexDjbqw_om.js";import{d as Q,a as W,b as X,c as Y,e as Z}from"./XdepartmentCncNztPP.js";import ee from"./XpremissionSetuvlfXZTk.js";import{_ as te}from"./X_plugin-vue_export-helperDlAUqK2U.js";const le={class:"main-content"},oe={class:"dialog-footer"},ne={__name:"department",setup(ae){const{proxy:v}=K(),g=c(0),p=c(!1),b=c([]),P=h([]);for(const[l]of Object.entries(T))P.push({key:l.__name,label:l.__name});const m=c(),o=h({id:0,pid:"",name:"",description:""}),U=h({pid:[{required:!0,trigger:"blur",validator:(l,e,d)=>{if(console.log(l,e,d),e===""){d(new Error("父级部门 不能为空"));return}if(e===o.id){d(new Error("所选部门不能为当前编辑部门"));return}d()}}],name:[{required:!0,message:"路由名称 不能为空",trigger:"blur"}]}),F=async l=>{l&&await l.validate((e,d)=>{if(!e)console.log("error submit!",d);else{let i={};console.log(o.id),o.id?i=Q(o):i=W(o),i.then(s=>{s.code===0&&(v.$message.success(s.msg),w(),p.value=!1,y(l))})}})},k=l=>{p.value=!1,y(l)},y=l=>{l&&l.resetFields()},C=(l=!1)=>{l?(o.id=l,X(l).then(e=>{e.code===0&&(o.pid=e.data.pid,o.name=e.data.name,o.description=e.data.description)})):o.id=0,E(),p.value=!0},B=l=>{l?Y(l).then(e=>{e.code===0&&(v.$message.success(e.msg),w())}):v.$message.error("请选择数据ID"),console.log("del",l)},w=()=>{Z().then(l=>{b.value=l.data})},V=c([{id:0,name:"根目录"}]),E=()=>{V.value=[{id:0,name:"根目录"}],$(b.value,V.value,!1)},$=(l,e,d)=>{l&&l.forEach(i=>{if(i.children&&i.children.length){const s={name:i.name,id:i.id,disabled:d||i.id===o.id,children:[]};$(i.children,s.children,d||i.id===o.id),e.push(s)}else{const s={name:i.name,id:i.id,disabled:d||i.id===o.id};e.push(s)}})};return L(()=>{w()}),(l,e)=>{const d=r("Plus"),i=r("el-icon"),s=r("el-button"),f=r("el-form-item"),_=r("el-table-column"),N=r("Edit"),O=r("Filter"),q=r("Delete"),z=r("el-popconfirm"),M=r("el-table"),S=r("el-cascader"),x=r("el-input"),j=r("el-form"),R=r("el-dialog");return I(),A("div",le,[t(f,{label:""},{default:n(()=>[t(s,{type:"primary",onClick:e[0]||(e[0]=a=>C())},{default:n(()=>[t(i,null,{default:n(()=>[t(d)]),_:1}),u(" 新建部门 ")]),_:1})]),_:1}),t(M,{data:b.value,"row-key":"id",height:"var(--global-table)",border:"","highlight-current-row":"","show-overflow-tooltip":"","default-expand-all":""},{default:n(()=>[t(_,{prop:"id",label:"ID",sortable:"",fixed:""}),t(_,{prop:"name",label:"部门名称"}),t(_,{prop:"description",label:"部门简介"}),t(_,{fixed:"right",label:"操作",width:"200"},{default:n(a=>[t(s,{link:"",type:"primary",size:"small",onClick:D=>C(a.row.id)},{default:n(()=>[t(i,null,{default:n(()=>[t(N)]),_:1}),u(" 编辑 ")]),_:2},1032,["onClick"]),t(s,{link:"",type:"primary",size:"small",onClick:D=>g.value=a.row.id},{default:n(()=>[t(i,null,{default:n(()=>[t(O)]),_:1}),u(" 权限设置 ")]),_:2},1032,["onClick"]),a.row.children?H("",!0):(I(),G(z,{key:0,title:"删除后不可恢复，确定删除菜单【"+a.row.name+"】?",onConfirm:D=>B(a.row.id),width:"200"},{reference:n(()=>[t(s,{link:"",type:"primary",size:"small"},{default:n(()=>[t(i,null,{default:n(()=>[t(q)]),_:1}),u(" 删除 ")]),_:1})]),_:2},1032,["title","onConfirm"]))]),_:1})]),_:1},8,["data"]),t(R,{modelValue:p.value,"onUpdate:modelValue":e[7]||(e[7]=a=>p.value=a),title:o.id?"编辑部门":"新建部门",width:"800",onClose:e[8]||(e[8]=a=>k(m.value))},{footer:n(()=>[J("div",oe,[t(s,{onClick:e[4]||(e[4]=a=>y(m.value))},{default:n(()=>[u("重置")]),_:1}),t(s,{onClick:e[5]||(e[5]=a=>k(m.value))},{default:n(()=>[u("取消")]),_:1}),t(s,{type:"primary",onClick:e[6]||(e[6]=a=>F(m.value))},{default:n(()=>[u("保存")]),_:1})])]),default:n(()=>[t(j,{ref_key:"formRef",ref:m,model:o,rules:U,"status-icon":"","label-position":"top","show-all-levels":!1,inline:""},{default:n(()=>[t(f,{label:"父级部门",prop:"pid",style:{width:"40%"}},{default:n(()=>[t(S,{modelValue:o.pid,"onUpdate:modelValue":e[1]||(e[1]=a=>o.pid=a),options:V.value,props:{expandTrigger:"hover",checkStrictly:!0,value:"id",label:"name",emitPath:!1},style:{width:"100%"}},null,8,["modelValue","options"])]),_:1}),t(f,{label:"部门名称",prop:"name",style:{width:"40%"}},{default:n(()=>[t(x,{modelValue:o.name,"onUpdate:modelValue":e[2]||(e[2]=a=>o.name=a)},null,8,["modelValue"])]),_:1}),t(f,{label:"部门简介",prop:"description",style:{width:"40%"}},{default:n(()=>[t(x,{modelValue:o.description,"onUpdate:modelValue":e[3]||(e[3]=a=>o.description=a)},null,8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])]),_:1},8,["modelValue","title"]),t(ee,{modelValue:g.value,"onUpdate:modelValue":e[9]||(e[9]=a=>g.value=a),type:"sys_department"},null,8,["modelValue"])])}}},ue=te(ne,[["__scopeId","data-v-9fb5b15e"]]);export{ue as default};
