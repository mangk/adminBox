import"./XvueCyDLjQZ5.js";import{p as e,b as l,c as t}from"./XindexXTMf5ase.js";import{_ as a}from"./X_plugin-vue_export-helperBCo6x5W8.js";import{bH as o,r as s,$ as i,f as r,ah as d,o as u,P as n,Q as p,a as m,X as c,V as f,U as v,W as y,c as _,F as j,a8 as w}from"./X@vueSdcp2x8a.js";import"./XpiniaDkaaaSbL.js";import"./Xelement-plusJnkMASPo.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCSz_Rkca.js";import"./X@element-plusCAKmCgXq.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./Xvue3-sfc-loaderDa78GGQm.js";import"./XaxiosCURSphCx.js";import"./XmittCNZ6avp8.js";import"./Xvue-routerCZMibURm.js";import"./XnprogressDgs0sPf-.js";const h={style:{display:"flex","flex-flow":"row nowrap","justify-content":"space-between"}},X={style:{"text-align":"left"}},b={style:{color:"#f56c6c","font-size":"12px"}},x={style:{color:"#e6a23c","font-size":"12px"}},k=a({__name:"premissionSet",props:{modelValue:{type:Number},modelModifiers:{},type:{type:String},typeModifiers:{}},emits:["update:modelValue","update:type"],setup(a){const k=o(a,"modelValue"),g=o(a,"type"),V=s(!1),O=s([]),z=i({}),P=s([]),C=i({});r(k,(t=>{t&&(V.value=!0,e().then((e=>{O.value=e.data})),l(k.value,g.value).then((e=>{var l,t,a,o,s;P.value=null==(l=e.data)?void 0:l.ohter_set_list,Object.assign(z,null==(a=null==(t=e.data)?void 0:t.cur_set)?void 0:a.list),Object.assign(C,null==(s=null==(o=e.data)?void 0:o.result_set)?void 0:s.list)})))})),r(z,(e=>{for(const l in e)Object.hasOwnProperty.call(e,l)&&(C[l]=U(l))}));const U=e=>{var l=0,t=[],a=[];z[e]&&(l=z[e]);for(const r in P.value)if(Object.hasOwnProperty.call(P.value,r)){const l=P.value[r];"sys_role"==l.module&&l.list&&l.list[e]&&t.push(l.list[e]),"sys_department"==l.module&&l.list&&l.list[e]&&a.push(l.list[e])}var o=F(t),s=F(a),i=0;return 0!==s&&(i=s),0!==o&&(i=o),0!==l&&(i=l),i},F=e=>{var l=0;for(let t=0;t<e.length;t++){const a=e[t];if(-1==a)return-1;1==a&&l++}return l>0?1:0},M=()=>{t({id:k.value,module:g.value,list:z}).then((()=>{I()}))},I=()=>{V.value=!1,k.value=0,P.value=[];for(const e in z)Object.hasOwnProperty.call(z,e)&&delete z[e];for(const e in C)Object.hasOwnProperty.call(C,e)&&delete C[e]};return(e,l)=>{const t=d("Menu"),a=d("Link"),o=d("Pointer"),s=d("el-icon"),i=d("el-table-column"),r=d("Check"),k=d("Close"),g=d("el-radio"),U=d("el-radio-group"),F=d("el-table"),S=d("el-button"),W=d("WarningFilled"),A=d("el-drawer");return u(),n(A,{modelValue:V.value,"onUpdate:modelValue":l[0]||(l[0]=e=>V.value=e),direction:"rtl","before-close":I,"with-header":!1,size:"90%"},{footer:p((()=>[m("div",h,[m("div",null,[c(S,{onClick:M,type:"primary"},{default:p((()=>[f("保存")])),_:1}),c(S,{onClick:I},{default:p((()=>[f("关闭")])),_:1})]),m("div",X,[m("div",b,[c(s,null,{default:p((()=>[c(W)])),_:1}),f(" 权限分级说明: 个人设置 > 角色设置 > 部门设置 ")]),m("div",x,[c(s,null,{default:p((()=>[c(W)])),_:1}),f(" 权限同级说明: 拒绝 > 允许; ")])])])])),default:p((()=>[c(F,{data:O.value,height:"calc(100vh - 62px - 40px )",style:{width:"100%"},"row-key":"id",border:"","default-expand-all":""},{default:p((()=>[c(i,{label:"菜单 · API · 动作",width:"190px",fixed:""},{default:p((e=>[c(s,null,{default:p((()=>["menu"==e.row.type?(u(),n(t,{key:0})):v("",!0),"api"==e.row.type?(u(),n(a,{key:1})):v("",!0),"action"==e.row.type?(u(),n(o,{key:2})):v("",!0)])),_:2},1024),f(" "+y(e.row.name),1)])),_:1}),c(i,{label:"结果展示",fixed:"",width:"80px"},{default:p((e=>[C[e.row.id]>0?(u(),n(s,{key:0,style:{color:"#67c23a"}},{default:p((()=>[c(r)])),_:1})):v("",!0),C[e.row.id]<0?(u(),n(s,{key:1,style:{color:"#f56c6c"}},{default:p((()=>[c(k)])),_:1})):v("",!0)])),_:1}),c(i,{label:"当前设置",fixed:"",width:"220px"},{default:p((e=>[c(U,{modelValue:z[e.row.id],"onUpdate:modelValue":l=>z[e.row.id]=l},{default:p((()=>[c(g,{size:"small",value:1},{default:p((()=>[f("允许")])),_:1}),c(g,{size:"small",value:-1},{default:p((()=>[f("拒绝")])),_:1}),c(g,{size:"small",value:0},{default:p((()=>[f("不限")])),_:1})])),_:2},1032,["modelValue","onUpdate:modelValue"])])),_:1}),(u(!0),_(j,null,w(P.value,((e,l)=>(u(),n(i,{width:"120px",key:l,label:e.record_name},{default:p((l=>[e.list&&e.list[l.row.id]&&e.list[l.row.id]>0?(u(),n(s,{key:0,style:{color:"#67c23a"}},{default:p((()=>[c(r)])),_:1})):v("",!0),e.list&&e.list[l.row.id]&&e.list[l.row.id]<0?(u(),n(s,{key:1,style:{color:"#f56c6c"}},{default:p((()=>[c(k)])),_:1})):v("",!0)])),_:2},1032,["label"])))),128))])),_:1},8,["data"])])),_:1},8,["modelValue"])}}},[["__scopeId","data-v-36096595"]]);export{k as default};