import{r as s,l as A,m as b,f as a,o,c as d,e as l,t as m,k as c,d as i,w as _,q as L,j as p,F as X,s as N}from"./XindexBzdqYFdI.js";import R from"./XmenuTreeCgbbnJC8.js";import{_ as D}from"./X_plugin-vue_export-helperDlAUqK2U.js";const S={class:"logo-box"},V=["src"],$={key:0,class:"logo-title"},q={key:0,class:"desc"},F={__name:"aside",async setup(O){let t,u;const v=s(window.adminX.Logo?window.adminX.Logo:"./images/logo.png"),f=s(window.adminX.Name),w=s(window.adminX.Desc?window.adminX.Desc:""),e=s(!1),g=([t,u]=A(()=>N().loadServerRouter()),t=await t,u(),t),h=(r,n)=>{},k=(r,n)=>{};return b(()=>{e.value=document.body.clientWidth<1200}),window.onresize=()=>(()=>{e.value=document.body.clientWidth<1200})(),(r,n)=>{const y=a("el-menu"),B=a("ArrowRightBold"),C=a("ArrowLeftBold"),x=a("el-icon");return o(),d(X,null,[l("div",null,[l("div",S,[l("img",{class:"logo",src:v.value,alt:""},null,8,V),e.value?c("",!0):(o(),d("div",$,m(f.value),1))]),e.value?c("",!0):(o(),d("div",q,m(w.value),1))]),i(y,{"default-active":r.$route.name,onOpen:h,onClose:k,collapse:e.value,"unique-opened":"",router:"","background-color":"#191a23","text-color":"#fff"},{default:_(()=>[i(R,{menus:L(g)},null,8,["menus"])]),_:1},8,["default-active","collapse"]),l("div",{onClick:n[0]||(n[0]=P=>e.value=!e.value),class:"collapse-change-btn"},[i(x,null,{default:_(()=>[e.value?(o(),p(B,{key:0})):c("",!0),e.value?c("",!0):(o(),p(C,{key:1}))]),_:1})])],64)}}},E=D(F,[["__scopeId","data-v-e2c99c23"]]);export{E as default};
