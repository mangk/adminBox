import"./XvueCyDLjQZ5.js";import{u as e,a as l}from"./Xindex4Ib4jN1z.js";import{f as a}from"./X@vueuseCSz_Rkca.js";import{_ as o}from"./X_plugin-vue_export-helperBCo6x5W8.js";import{$ as t,r as s,bP as r,H as u,ah as n,o as i,P as d,Q as c,X as m,a as v,c as p,W as f,U as g,u as y,a8 as b,F as h,V as x,L as j,Z as k,aI as X,S as _}from"./X@vueSdcp2x8a.js";import"./XpiniaDkaaaSbL.js";import"./Xelement-plusViPJHclL.js";import"./Xlodash-esBg5u8xPa.js";import"./X@element-plusCAKmCgXq.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./XaxiosCURSphCx.js";import"./XmittCNZ6avp8.js";import"./Xvue-routerCZMibURm.js";import"./XnprogressDgs0sPf-.js";import"./Xvue3-sfc-loaderDa78GGQm.js";const w=["src"],S={key:0,class:"header-name"},C={class:"left header-user"},I={key:0},O={class:"demo-rich-conent",style:{display:"flex",gap:"6px","flex-direction":"column"}},V={key:0},z={style:{display:"flex","justify-content":"space-between"}},N={style:{display:"flex","justify-content":"space-between"}},P="#2d2d32",B=o({__name:"styleDefault",async setup(o){let B,J;const U=t(localStorage.getItem("x-theme-config")?JSON.parse(localStorage.getItem("x-theme-config")):{darkSidebar:!0}),$=s(!1),F=s(U.headerMenu),L=e=>{U.headerMenu=e,localStorage.setItem("x-theme-config",JSON.stringify(U))},M=s(U.darkSidebar),W=e=>{U.darkSidebar=e,localStorage.setItem("x-theme-config",JSON.stringify(U))},q=a("--el-color-primary");U.colorPrimary&&(q.value=U.colorPrimary);const D=[{value:"#E63415",label:"red"},{value:"#e0620e",label:"orange"},{value:"#1EC79D",label:"green"},{value:"#4167F0",label:"blue"},{value:"#6222C9",label:"purple"},{value:"#000",label:"black"}],E=s(window.adminBox.Logo?window.adminBox.Logo:"./images/logo.svg"),H=s(window.adminBox.Name),Q=e(),R=s({});R.value=([B,J]=r((()=>Q.userInfo())),B=await B,J(),B);const T=()=>{Q.logOut()},Z=s(!1),A=([B,J]=r((()=>l().loadServerRouter())),B=await B,J(),B),G=(e,l)=>{},K=(e,l)=>{},Y=()=>{document.body.clientWidth>=1100?Z.value=!1:Z.value=!0,$.value=document.body.clientWidth<=550};return u(Y),window.onresize=()=>Y(),(e,l)=>{const a=n("MenuTree"),o=n("el-menu"),t=n("el-scrollbar"),s=n("el-avatar"),r=n("el-tag"),u=n("el-switch"),B=n("el-collapse-item"),J=n("el-collapse"),Q=n("el-button"),Y=n("el-popover"),ee=n("el-header"),le=n("el-aside"),ae=n("router-view"),oe=n("el-main"),te=n("el-container");return i(),d(te,{style:{width:"100%",height:"100%"}},{default:c((()=>[m(ee,{class:"box-header",style:j({"background-color":M.value?P:"",color:M.value?"#fff":""})},{default:c((()=>[v("img",{class:"header-logo",src:E.value},null,8,w),$.value?g("",!0):(i(),p("div",S,f(H.value),1)),m(t,{style:{height:"unset",margin:"0 8px"}},{default:c((()=>[F.value||$.value?(i(),d(o,{key:0,class:"header-menu","default-active":e.$route.name,onOpen:G,onClose:K,collapse:Z.value,"unique-opened":"",router:"",mode:"horizontal",ellipsis:!1,"background-color":M.value?P:"","text-color":M.value?"#fff":"",style:{width:"auto"}},{default:c((()=>[m(a,{menus:y(A)},null,8,["menus"])])),_:1},8,["default-active","collapse","background-color","text-color"])):g("",!0)])),_:1}),m(Y,null,{reference:c((()=>[v("div",C,[$.value?g("",!0):(i(),p("span",I,f(R.value.username),1)),m(s,{src:R.value.avatar,icon:"UserFilled",size:30,style:{"margin-left":"5px"}},null,8,["src"])])])),default:c((()=>[v("div",O,[$.value?(i(),p("span",V,f(R.value.username),1)):g("",!0),m(J,{accordion:""},{default:c((()=>[m(B,{title:"主题设置",name:"1"},{default:c((()=>[v("div",null,[(i(),p(h,null,b(D,(e=>m(r,{key:e,color:e.value,onClick:l=>{return a=e.value,q.value=a,U.colorPrimary=a,void localStorage.setItem("x-theme-config",JSON.stringify(U));var a}},null,8,["color","onClick"]))),64))]),v("div",z,[x(" 顶部菜单"),m(u,{modelValue:F.value,"onUpdate:modelValue":l[0]||(l[0]=e=>F.value=e),onChange:L,disabled:$.value},null,8,["modelValue","disabled"])]),v("div",N,[x(" 深色边栏"),m(u,{modelValue:M.value,"onUpdate:modelValue":l[1]||(l[1]=e=>M.value=e),onChange:W},null,8,["modelValue"])])])),_:1})])),_:1}),m(Q,{link:"",onClick:T},{default:c((()=>[x("退出")])),_:1})])])),_:1})])),_:1},8,["style"]),m(te,{class:"box-asside-and-main"},{default:c((()=>[F.value||$.value?g("",!0):(i(),d(le,{key:0,class:"box-aside",style:j({"background-color":M.value?P:"",color:M.value?"#fff":""})},{default:c((()=>[m(t,null,{default:c((()=>[m(k,{duration:{enter:800,leave:100},mode:"out-in",name:"el-fade-in-linear"},{default:c((()=>[m(o,{"default-active":e.$route.name,onOpen:G,onClose:K,collapse:Z.value,"unique-opened":"",router:"","background-color":M.value?P:"","text-color":M.value?"#fff":""},{default:c((()=>[m(a,{menus:y(A)},null,8,["menus"])])),_:1},8,["default-active","collapse","background-color","text-color"])])),_:1})])),_:1})])),_:1},8,["style"])),m(oe,{class:"box-main"},{default:c((()=>[m(ae,null,{default:c((({Component:l})=>[(i(),d(X,null,[(i(),d(_(l),{key:e.$route.path}))],1024))])),_:1})])),_:1})])),_:1})])),_:1})}}},[["__scopeId","data-v-186b7e99"]]);export{B as default};
