const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["./X404DLzpbq3O.js","./X_plugin-vue_export-helperBCo6x5W8.js","./XloginDTkozuwP.js","./XvueBaz_cZdj.js","./X@vuen9TiMG2O.js","./Xvue-routerzsXVQF6n.js","./XpiniaYeQOi8UL.js","./Xelement-plusCLM398fa.js","./Xlodash-esBg5u8xPa.js","./X@vueuseCCX6pF6_.js","./X@element-plusBnJF7_W-.js","./X@popperjsD3lHDW-0.js","./X@ctrlD2oWfImC.js","./Xdayjst441MSVg.js","./Xasync-validatorCuo4gI4y.js","./Xmemoize-oneDs0C_khL.js","./Xnormalize-wheel-esVn5vHDCm.js","./X@floating-uipMauM7H8.js","./XaxiosByKc3Ecy.js","./XmittCNZ6avp8.js","./XnprogressDgs0sPf-.js","./Xvue3-sfc-loaderDa78GGQm.js","./XmenuTreeDJLeaEga.js","./XstyleDefaultB95cqfrl.js","./XtabsADncIrG.js","./XapiBbj6_A9d.js","./XmenuD1RrOB4o.js","./XdepartmentBGBGGnAp.js","./XdepartmentgA6gy3mg.js","./XpremissionSetB7hsfEEY.js","./XindexCqpEmpml.js","./XfileUploadCBybYkC5.js","./XmenucJmI6mAr.js","./XroleB-pRVrki.js","./XroleCcdS7NM9.js","./XuserB07kmE2D.js","./XserverComponentC8510u9D.js","./XwelcomeC41FH4wA.js"])))=>i.map(i=>d[i]);
import{V as e}from"./XvueBaz_cZdj.js";import{d as t,c as o}from"./XpiniaYeQOi8UL.js";import{E as r,a as n,z as a,b as i,c as s}from"./Xelement-plusCLM398fa.js";import{Q as l}from"./X@element-plusBnJF7_W-.js";import{a as u}from"./XaxiosByKc3Ecy.js";import{m}from"./XmittCNZ6avp8.js";import{c,a as d}from"./Xvue-routerzsXVQF6n.js";import{N as p}from"./XnprogressDgs0sPf-.js";import{b as f,am as _,aY as g,ah as v,o as h,P as w,Q as y,aL as E,X as j,u as T,as as P}from"./X@vuen9TiMG2O.js";import{i as X}from"./Xvue3-sfc-loaderDa78GGQm.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCCX6pF6_.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver((e=>{for(const o of e)if("childList"===o.type)for(const e of o.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&t(e)})).observe(document,{childList:!0,subtree:!0})}function t(e){if(e.ep)return;e.ep=!0;const t=function(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerPolicy&&(t.referrerPolicy=e.referrerPolicy),"use-credentials"===e.crossOrigin?t.credentials="include":"anonymous"===e.crossOrigin?t.credentials="omit":t.credentials="same-origin",t}(e);fetch(e.href,t)}}();const I=m(),R={},L=function(e,t,o){let r=Promise.resolve();if(t&&t.length>0){const e=document.getElementsByTagName("link"),n=document.querySelector("meta[property=csp-nonce]"),a=(null==n?void 0:n.nonce)||(null==n?void 0:n.getAttribute("nonce"));r=Promise.allSettled(t.map((t=>{if(t=function(e,t){return new URL(e,t).href}(t,o),t in R)return;R[t]=!0;const r=t.endsWith(".css"),n=r?'[rel="stylesheet"]':"";if(!!o)for(let o=e.length-1;o>=0;o--){const n=e[o];if(n.href===t&&(!r||"stylesheet"===n.rel))return}else if(document.querySelector(`link[href="${t}"]${n}`))return;const i=document.createElement("link");return i.rel=r?"stylesheet":"modulepreload",r||(i.as="script"),i.crossOrigin="",i.href=t,a&&i.setAttribute("nonce",a),document.head.appendChild(i),r?new Promise(((e,o)=>{i.addEventListener("load",e),i.addEventListener("error",(()=>o(new Error(`Unable to preload CSS for ${t}`))))})):void 0})))}function n(e){const t=new Event("vite:preloadError",{cancelable:!0});if(t.payload=e,window.dispatchEvent(t),!t.defaultPrevented)throw e}return r.then((t=>{for(const e of t||[])"rejected"===e.status&&n(e.reason);return e().catch(n)}))},A=()=>J({url:"/sys/auth/permissionAll",method:"GET"}),x=(e,t)=>J({url:"/sys/auth/permissionGetByIdAndModule",method:"POST",data:{id:e,module:t}}),S=e=>J({url:"/sys/auth/permissionSave",method:"PUT",data:e}),O=()=>J({url:"/sys/verificationCode",method:"GET"}),b=Object.assign({"../views/404.vue":()=>L((()=>import("./X404DLzpbq3O.js")),__vite__mapDeps([0,1]),import.meta.url),"../views/login.vue":()=>L((()=>import("./XloginDTkozuwP.js")),__vite__mapDeps([2,3,4,5,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]),import.meta.url),"../views/main/menuTree.vue":()=>L((()=>import("./XmenuTreeDJLeaEga.js")),__vite__mapDeps([22,3,4]),import.meta.url),"../views/main/styleDefault.vue":()=>L((()=>import("./XstyleDefaultB95cqfrl.js")),__vite__mapDeps([23,3,4,9,6,7,8,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/main/tab.vue":()=>L((()=>import("./XtabsADncIrG.js")),__vite__mapDeps([24,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]),import.meta.url),"../views/setting/api.vue":()=>L((()=>import("./XapiBbj6_A9d.js")),__vite__mapDeps([25,3,4,26,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/department.vue":()=>L((()=>import("./XdepartmentBGBGGnAp.js")),__vite__mapDeps([27,3,4,10,28,29,1,6,7,8,9,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/index.vue":()=>L((()=>import("./XindexCqpEmpml.js")),__vite__mapDeps([30,3,4,31,7,8,9,10,11,12,13,14,15,16,17,6,18,19,5,20,21]),import.meta.url),"../views/setting/menu.vue":()=>L((()=>import("./XmenucJmI6mAr.js")),__vite__mapDeps([32,3,4,26,10,1,6,7,8,9,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/premissionSet.vue":()=>L((()=>import("./XpremissionSetB7hsfEEY.js")),__vite__mapDeps([29,3,4,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/role.vue":()=>L((()=>import("./XroleB-pRVrki.js")),__vite__mapDeps([33,3,4,34,29,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/user.vue":()=>L((()=>import("./XuserB07kmE2D.js")),__vite__mapDeps([35,3,4,29,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21,28,34,31]),import.meta.url),"../views/util/fileUpload.vue":()=>L((()=>import("./XfileUploadCBybYkC5.js")),__vite__mapDeps([31,3,4,7,8,9,10,11,12,13,14,15,16,17,6,18,19,5,20,21]),import.meta.url),"../views/util/serverComponent.vue":()=>L((()=>import("./XserverComponentC8510u9D.js")),__vite__mapDeps([36,3,4,5,1]),import.meta.url),"../views/welcome.vue":()=>L((()=>import("./XwelcomeC41FH4wA.js")),__vite__mapDeps([37,3,4,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url)}),B=Object.assign({});function D(e){const t=Object.keys(b).filter((t=>t.replace("../","")===e))[0];return b[t]}function V(e){e.length&&e.forEach((e=>{e.component&&("views"===e.component.split("/")[0]?e.component=D(e.component):"plugins"===e.component.split("/")[0]&&(e.component=function(e){const t=Object.keys(B).filter((t=>t.replace("../","")===e))[0];return B[t]}(e.component))),e.children&&e.children.length&&(delete e.component,V(e.children))}))}function k(){var e="admin";return window._adminBox&&window._adminBox.backgroundPrefix&&(e=window._adminBox.backgroundPrefix),e}const C=(e=!1)=>{var t="";-1==["/","/login"].indexOf(H.currentRoute.value.path)&&(r({showClose:!0,message:"身份认证失败",type:"error"}),t=encodeURIComponent(H.currentRoute.value.fullPath)),H.currentRoute.value.query.redirect&&(t=H.currentRoute.value.query.redirect);let o={path:"/login"};t&&(o.query={redirect:t}),H.replace(o)},U=t("user",(()=>{const e=f(0),t=f({}),o=()=>(window.adminBox.Name?window.adminBox.Name+"-":"")+"x-token",n=()=>(window.adminBox.Name?window.adminBox.Name+"-":"")+"x-user-type",a=()=>(window.adminBox.Name?window.adminBox.Name+"-":"")+"x-user-id";return{initialized:e,userInfo:(o=!1)=>((o||0===e.value)&&(e.value=1),t.value),setUserData:e=>{t.value=e},logIn:async i=>{const s=k();var l,u=await(l=i,J({url:"/sys/login",method:"POST",data:l}));if(0!=u.code)return r.error(u.msg),"";t.value=u.data,e.value=1,await localStorage.setItem(o(),u.data.jwt_token),await localStorage.setItem(n(),"default"),await localStorage.setItem(a(),u.data.id);const m=F();return await m.loadServerRouter(!0),H.currentRoute.value.query.redirect?decodeURIComponent(H.currentRoute.value.query.redirect):u.data.user_config.home_page?u.data.user_config.home_page:`/${s}/welcome`},logOut:()=>{localStorage.clear(),window.location.href=window.location.origin},isLogIn:()=>localStorage.getItem(o())&&localStorage.getItem(n())&&localStorage.getItem(a()),userAuth:()=>{let e=localStorage.getItem(o()),t=localStorage.getItem(n()),r=localStorage.getItem(a());return e&&t&&r?{token:e,userType:t,userId:r}:(C(!0),!1)}}})),N={};function $(e,t=null){if(N[e])return void("function"==typeof t&&(N[e].push(t),t()));N[e]=[],"function"==typeof t&&N[e].push(t);const o=document.createElement("script");o.type="text/javascript",o.src=e,o.onload=()=>{N[e].forEach((e=>e()))},document.head.appendChild(o)}function q(t,o="myConvert"){const r={moduleCache:{vue:e},getFile:()=>J(t,{method:"GET"}).then((e=>e)),addStyle(e){const t=document.createElement("style");t.setAttribute("id",o),t.textContent=e;const r=document.head.getElementsByTagName("style")[document.head.getElementsByTagName("style").length-1]||null;document.head.insertBefore(t,r)}},n=g((()=>X(`${o}.vue`,r)));return _(n)}const z=k(),G={default:D("views/main/styleDefault.vue")},M=e=>(G[e]||(G[e]=q(e)),G[e]),F=t("router",(()=>{const e=f(0),t=f([]);return{initialized:e,loadServerRouter:async(o=!1)=>{const r=U();if(o||0===e.value){e.value=1;const o=await J({url:"/sys/auth/userPermission",method:"GET"});t.value=o.data.menu[0].children,r.setUserData(o.data.user),V(t.value);let n=window.adminBox.Theme;H.addRoute({path:"/"+z,name:z,meta:{icon:"add"},component:M(n||"default"),children:t.value})}return t.value}}})),Q=k(),H=c({history:d("./"),routes:[{path:"/",redirect:`/${Q}/welcome`},{path:"/login",name:"login",component:()=>L((()=>import("./XloginDTkozuwP.js")),__vite__mapDeps([2,3,4,5,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]),import.meta.url)},{path:"/"+Q,name:Q,component:()=>L((()=>import("./XstyleDefaultB95cqfrl.js")),__vite__mapDeps([23,3,4,9,6,7,8,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url)},{path:"/:catchAll(.*)",meta:{closeTab:!0},component:()=>L((()=>import("./X404DLzpbq3O.js")),__vite__mapDeps([0,1]),import.meta.url)}]});H.beforeEach((async(e,t)=>{if(p.start(),e.meta.title&&(window.document.title=e.meta.title),"/login"===e.path)return!0;const o=F(),r=U();return o.initialized?(r.isLogIn()||r.logOut(),!0):(await o.loadServerRouter(!0),{path:e.path})})),H.afterEach((()=>{p.done()})),H.onError((()=>{p.done()}));const J=u.create({baseURL:Z(),timeout:3e4});let K,W=0;const Y=()=>{W--,W<=0&&(clearTimeout(K),I.emit("closeLoading"))};function Z(){let e=[];return window.adminBox.RunAt&&e.push(window.adminBox.RunAt),window.adminBox.BackendRouterPrefix&&e.push(window.adminBox.BackendRouterPrefix),e.join("/")}J.interceptors.request.use((e=>{e.donNotShowLoading||(W++,K&&clearTimeout(K),K=setTimeout((()=>{W>0&&I.emit("showLoading")}),400));const t=U().userAuth();return e.headers={"Content-Type":"application/json; charset=utf-8",Authorization:"Bearer "+t.token,"X-User-Type":t.userType,"X-User-Id":t.userId,...e.headers},e}),(e=>(e.config.donNotShowLoading||Y(),r({showClose:!0,message:e,type:"error"}),e))),J.interceptors.response.use((e=>e.data),(e=>{if(Y(),401===e)C(!0);else n.confirm(`\n          <p>检测到接口错误${e}</p>\n          <p>错误码<span style="color:red"> ${e.response.status} </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>\n          `,"接口报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"我知道了",cancelButtonText:"取消"});return e}));const ee={install:e=>{(async e=>{for(const r in l)e.component(r,l[r]);const t=await D("views/util/fileUpload.vue")(),o=await D("views/main/menuTree.vue")();e.component("FileUpload",t.default||t),e.component("MenuTree",o.default||o),e.config.globalProperties.$loadJS=$,e.config.globalProperties.$loadTMPL=q,e.config.globalProperties.$message=r,e.config.globalProperties.$http=J,e.config.globalProperties.$useUserStore=U,e.config.globalProperties.$useRouterStore=F})(e)}},te={__name:"App",setup(e){let t="";if("en"===window.adminBox.Locale)t=i;else t=a;return(e,o)=>{const r=v("router-view"),n=v("el-config-provider");return h(),w(n,{locale:T(t)},{default:y((()=>[(h(),w(E,null,{default:y((()=>[j(r)])),_:1}))])),_:1},8,["locale"])}}};p.configure({showSpinner:!1,ease:"ease",speed:200});const oe=P(te);oe.use(o()),oe.use(s),oe.use(ee),oe.use(H),oe.mount("#app");export{F as a,x as b,S as c,I as e,J as h,A as p,Z as s,U as u,O as v};
