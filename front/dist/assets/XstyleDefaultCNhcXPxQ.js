import"./XvueCyDLjQZ5.js";import e from"./XmenuTreeDxIhcR46.js";import{u as l,a}from"./XindexDacNJLVO.js";import{f as o}from"./X@vueuseCSz_Rkca.js";import{_ as t}from"./X_plugin-vue_export-helperBCo6x5W8.js";import{$ as s,r,bP as u,H as n,ah as i,o as d,P as c,Q as m,X as v,a as p,c as f,W as g,U as y,u as b,a8 as h,F as j,V as x,L as X,Z as k,aI as _,S as w}from"./X@vueSdcp2x8a.js";import"./XpiniaDkaaaSbL.js";import"./Xelement-plusViPJHclL.js";import"./Xlodash-esBg5u8xPa.js";import"./X@element-plusCAKmCgXq.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./Xvue3-sfc-loaderDa78GGQm.js";import"./XaxiosCURSphCx.js";import"./XmittCNZ6avp8.js";import"./Xvue-routerCZMibURm.js";import"./XnprogressDgs0sPf-.js";const S=["src"],C={key:0,class:"header-name"},I={class:"left header-user"},O={key:0},V={class:"demo-rich-conent",style:{display:"flex",gap:"6px","flex-direction":"column"}},z={key:0},N={style:{display:"flex","justify-content":"space-between"}},P={style:{display:"flex","justify-content":"space-between"}},F="#2d2d32",J=t({__name:"styleDefault",async setup(t){let J,U;const $=s(localStorage.getItem("x-theme-config")?JSON.parse(localStorage.getItem("x-theme-config")):{darkSidebar:!0}),B=r(!1),L=r($.headerMenu),W=e=>{$.headerMenu=e,localStorage.setItem("x-theme-config",JSON.stringify($))},q=r($.darkSidebar),D=e=>{$.darkSidebar=e,localStorage.setItem("x-theme-config",JSON.stringify($))},E=o("--el-color-primary");$.colorPrimary&&(E.value=$.colorPrimary);const M=[{value:"#E63415",label:"red"},{value:"#e0620e",label:"orange"},{value:"#1EC79D",label:"green"},{value:"#4167F0",label:"blue"},{value:"#6222C9",label:"purple"},{value:"#000",label:"black"}],H=r(window.adminBox.Logo?window.adminBox.Logo:"./images/logo.svg"),Q=r(window.adminBox.Name),R=l(),T=r({});T.value=([J,U]=u((()=>R.userInfo())),J=await J,U(),J);const Z=()=>{R.logOut()},A=r(!1),G=([J,U]=u((()=>a().loadServerRouter())),J=await J,U(),J),K=(e,l)=>{},Y=(e,l)=>{},ee=()=>{document.body.clientWidth>=1100?A.value=!1:A.value=!0,B.value=document.body.clientWidth<=550};return n(ee),window.onresize=()=>ee(),(l,a)=>{const o=i("el-menu"),t=i("el-scrollbar"),s=i("el-avatar"),r=i("el-tag"),u=i("el-switch"),n=i("el-collapse-item"),J=i("el-collapse"),U=i("el-button"),R=i("el-popover"),ee=i("el-header"),le=i("el-aside"),ae=i("router-view"),oe=i("el-main"),te=i("el-container");return d(),c(te,{style:{width:"100%",height:"100%"}},{default:m((()=>[v(ee,{class:"box-header",style:X({"background-color":q.value?F:"",color:q.value?"#fff":""})},{default:m((()=>[p("img",{class:"header-logo",src:H.value},null,8,S),B.value?y("",!0):(d(),f("div",C,g(Q.value),1)),v(t,{style:{height:"unset",margin:"0 8px"}},{default:m((()=>[L.value||B.value?(d(),c(o,{key:0,class:"header-menu","default-active":l.$route.name,onOpen:K,onClose:Y,collapse:A.value,"unique-opened":"",router:"",mode:"horizontal",ellipsis:!1,"background-color":q.value?F:"","text-color":q.value?"#fff":"",style:{width:"auto"}},{default:m((()=>[v(e,{menus:b(G)},null,8,["menus"])])),_:1},8,["default-active","collapse","background-color","text-color"])):y("",!0)])),_:1}),v(R,null,{reference:m((()=>[p("div",I,[B.value?y("",!0):(d(),f("span",O,g(T.value.username),1)),v(s,{src:T.value.avatar,icon:"UserFilled",size:30,style:{"margin-left":"5px"}},null,8,["src"])])])),default:m((()=>[p("div",V,[B.value?(d(),f("span",z,g(T.value.username),1)):y("",!0),v(J,{accordion:""},{default:m((()=>[v(n,{title:"主题设置",name:"1"},{default:m((()=>[p("div",null,[(d(),f(j,null,h(M,(e=>v(r,{key:e,color:e.value,onClick:l=>{return a=e.value,E.value=a,$.colorPrimary=a,void localStorage.setItem("x-theme-config",JSON.stringify($));var a}},null,8,["color","onClick"]))),64))]),p("div",N,[x(" 顶部菜单"),v(u,{modelValue:L.value,"onUpdate:modelValue":a[0]||(a[0]=e=>L.value=e),onChange:W,disabled:B.value},null,8,["modelValue","disabled"])]),p("div",P,[x(" 深色边栏"),v(u,{modelValue:q.value,"onUpdate:modelValue":a[1]||(a[1]=e=>q.value=e),onChange:D},null,8,["modelValue"])])])),_:1})])),_:1}),v(U,{link:"",onClick:Z},{default:m((()=>[x("退出")])),_:1})])])),_:1})])),_:1},8,["style"]),v(te,{class:"box-asside-and-main"},{default:m((()=>[L.value||B.value?y("",!0):(d(),c(le,{key:0,class:"box-aside",style:X({"background-color":q.value?F:"",color:q.value?"#fff":""})},{default:m((()=>[v(t,null,{default:m((()=>[v(k,{duration:{enter:800,leave:100},mode:"out-in",name:"el-fade-in-linear"},{default:m((()=>[v(o,{"default-active":l.$route.name,onOpen:K,onClose:Y,collapse:A.value,"unique-opened":"",router:"","background-color":q.value?F:"","text-color":q.value?"#fff":""},{default:m((()=>[v(e,{menus:b(G)},null,8,["menus"])])),_:1},8,["default-active","collapse","background-color","text-color"])])),_:1})])),_:1})])),_:1},8,["style"])),v(oe,{class:"box-main"},{default:m((()=>[v(ae,null,{default:m((({Component:e})=>[(d(),c(_,null,[(d(),c(w(e),{key:l.$route.path}))],1024))])),_:1})])),_:1})])),_:1})])),_:1})}}},[["__scopeId","data-v-81bd1fc4"]]);export{J as default};
