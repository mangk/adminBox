const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["./X404DLzpbq3O.js","./X_plugin-vue_export-helperBCo6x5W8.js","./XloginCqK-bZSU.js","./XvueBaz_cZdj.js","./X@vuen9TiMG2O.js","./Xvue-routerzsXVQF6n.js","./XpiniaYeQOi8UL.js","./Xelement-plusCLM398fa.js","./Xlodash-esBg5u8xPa.js","./X@vueuseCCX6pF6_.js","./X@element-plusBnJF7_W-.js","./X@popperjsD3lHDW-0.js","./X@ctrlD2oWfImC.js","./Xdayjst441MSVg.js","./Xasync-validatorCuo4gI4y.js","./Xmemoize-oneDs0C_khL.js","./Xnormalize-wheel-esVn5vHDCm.js","./X@floating-uipMauM7H8.js","./XaxiosByKc3Ecy.js","./XmittCNZ6avp8.js","./XnprogressDgs0sPf-.js","./Xvue3-sfc-loaderDa78GGQm.js","./XmenuTreeDJLeaEga.js","./XstyleDefaultB4IU6vd8.js","./XtabDhbf9wJD.js","./XapiBfDNUOpF.js","./XmenuBx33FT-3.js","./XdepartmentmKls9tzH.js","./XdepartmentDDNHFgsn.js","./XpremissionSetv_mz_vkC.js","./XindexDdrtUNH3.js","./XmenuDY1MwHPZ.js","./Xrolen37daFo-.js","./XroleDGeLKDdZ.js","./XuserD86X-_g6.js","./XfileUploadBKLVqaMF.js","./XserverComponentC8510u9D.js","./XwelcomeBIaRi45K.js"])))=>i.map(i=>d[i]);
import{V as e}from"./XvueBaz_cZdj.js";import{d as t,c as o}from"./XpiniaYeQOi8UL.js";import{E as r,a as n,z as s,b as a,c as i}from"./Xelement-plusCLM398fa.js";import{Q as l}from"./X@element-plusBnJF7_W-.js";import{a as u}from"./XaxiosByKc3Ecy.js";import{m}from"./XmittCNZ6avp8.js";import{c,a as p}from"./Xvue-routerzsXVQF6n.js";import{N as d}from"./XnprogressDgs0sPf-.js";import{b as f,am as g,aY as _,ah as h,o as v,P as w,Q as y,aL as E,X as T,u as j,as as P}from"./X@vuen9TiMG2O.js";import{i as L}from"./Xvue3-sfc-loaderDa78GGQm.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCCX6pF6_.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver((e=>{for(const o of e)if("childList"===o.type)for(const e of o.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&t(e)})).observe(document,{childList:!0,subtree:!0})}function t(e){if(e.ep)return;e.ep=!0;const t=function(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerPolicy&&(t.referrerPolicy=e.referrerPolicy),"use-credentials"===e.crossOrigin?t.credentials="include":"anonymous"===e.crossOrigin?t.credentials="omit":t.credentials="same-origin",t}(e);fetch(e.href,t)}}();const X=m(),I={},S=function(e,t,o){let r=Promise.resolve();if(t&&t.length>0){const e=document.getElementsByTagName("link"),n=document.querySelector("meta[property=csp-nonce]"),s=(null==n?void 0:n.nonce)||(null==n?void 0:n.getAttribute("nonce"));r=Promise.allSettled(t.map((t=>{if(t=function(e,t){return new URL(e,t).href}(t,o),t in I)return;I[t]=!0;const r=t.endsWith(".css"),n=r?'[rel="stylesheet"]':"";if(!!o)for(let o=e.length-1;o>=0;o--){const n=e[o];if(n.href===t&&(!r||"stylesheet"===n.rel))return}else if(document.querySelector(`link[href="${t}"]${n}`))return;const a=document.createElement("link");return a.rel=r?"stylesheet":"modulepreload",r||(a.as="script"),a.crossOrigin="",a.href=t,s&&a.setAttribute("nonce",s),document.head.appendChild(a),r?new Promise(((e,o)=>{a.addEventListener("load",e),a.addEventListener("error",(()=>o(new Error(`Unable to preload CSS for ${t}`))))})):void 0})))}function n(e){const t=new Event("vite:preloadError",{cancelable:!0});if(t.payload=e,window.dispatchEvent(t),!t.defaultPrevented)throw e}return r.then((t=>{for(const e of t||[])"rejected"===e.status&&n(e.reason);return e().catch(n)}))},b=()=>F({url:"/sys/auth/permissionAll",method:"GET"}),A=(e,t)=>F({url:"/sys/auth/permissionGetByIdAndModule",method:"POST",data:{id:e,module:t}}),O=e=>F({url:"/sys/auth/permissionSave",method:"PUT",data:e}),x=()=>F({url:"/sys/verificationCode",method:"GET"}),R=Object.assign({"../views/404.vue":()=>S((()=>import("./X404DLzpbq3O.js")),__vite__mapDeps([0,1]),import.meta.url),"../views/login.vue":()=>S((()=>import("./XloginCqK-bZSU.js")),__vite__mapDeps([2,3,4,5,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]),import.meta.url),"../views/main/menuTree.vue":()=>S((()=>import("./XmenuTreeDJLeaEga.js")),__vite__mapDeps([22,3,4]),import.meta.url),"../views/main/styleDefault.vue":()=>S((()=>import("./XstyleDefaultB4IU6vd8.js")),__vite__mapDeps([23,3,4,9,1,6,7,8,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/main/tab.vue":()=>S((()=>import("./XtabDhbf9wJD.js")),__vite__mapDeps([24,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]),import.meta.url),"../views/setting/api.vue":()=>S((()=>import("./XapiBfDNUOpF.js")),__vite__mapDeps([25,3,4,26,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/department.vue":()=>S((()=>import("./XdepartmentmKls9tzH.js")),__vite__mapDeps([27,3,4,10,28,29,1,6,7,8,9,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/index.vue":()=>S((()=>import("./XindexDdrtUNH3.js")),__vite__mapDeps([30,1]),import.meta.url),"../views/setting/menu.vue":()=>S((()=>import("./XmenuDY1MwHPZ.js")),__vite__mapDeps([31,3,4,26,10,1,6,7,8,9,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/premissionSet.vue":()=>S((()=>import("./XpremissionSetv_mz_vkC.js")),__vite__mapDeps([29,3,4,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/role.vue":()=>S((()=>import("./Xrolen37daFo-.js")),__vite__mapDeps([32,3,4,33,29,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url),"../views/setting/user.vue":()=>S((()=>import("./XuserD86X-_g6.js")),__vite__mapDeps([34,3,4,29,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21,28,33,35]),import.meta.url),"../views/util/fileUpload.vue":()=>S((()=>import("./XfileUploadBKLVqaMF.js")),__vite__mapDeps([35,3,4,7,8,9,10,11,12,13,14,15,16,17,6,18,19,5,20,21]),import.meta.url),"../views/util/serverComponent.vue":()=>S((()=>import("./XserverComponentC8510u9D.js")),__vite__mapDeps([36,3,4,5,1]),import.meta.url),"../views/welcome.vue":()=>S((()=>import("./XwelcomeBIaRi45K.js")),__vite__mapDeps([37,3,4,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url)}),D=Object.assign({});function B(e){const t=Object.keys(R).filter((t=>t.replace("../","")===e))[0];return R[t]}function V(e){e.length&&e.forEach((e=>{e.component&&("views"===e.component.split("/")[0]?e.component=B(e.component):"plugins"===e.component.split("/")[0]&&(e.component=function(e){const t=Object.keys(D).filter((t=>t.replace("../","")===e))[0];return D[t]}(e.component))),e.children&&e.children.length&&(delete e.component,V(e.children))}))}function k(){var e="admin";return window._adminBox&&window._adminBox.backgroundPrefix&&(e=window._adminBox.backgroundPrefix),e}const C=t("user",(()=>{const e=f(0),t=f({}),o="x-token",r="x-user-type",n="x-user-id";return{initialized:e,userInfo:(o=!1)=>((o||0===e.value)&&(e.value=1),t.value),setUserData:e=>{t.value=e},token:()=>{var e=localStorage.getItem(o);if(e)return e;H.replace({name:"login"})},userType:()=>localStorage.getItem(r),userId:()=>localStorage.getItem(n),logIn:async s=>{const a=k();var i,l=await(i=s,F({url:"/sys/login",method:"POST",data:i}));if(0!=l.code)return"";t.value=l.data,e.value=1,await localStorage.setItem(o,l.data.jwt_token),await localStorage.setItem(r,"default"),await localStorage.setItem(n,l.data.id);const u=q();return await u.loadServerRouter(!0),l.data.user_config.home_page?l.data.user_config.home_page:`/${a}/welcome`},logOut:()=>{localStorage.clear(),window.location.href=window.location.origin},isLogIn:()=>localStorage.getItem(o)&&localStorage.getItem(r)&&localStorage.getItem(n)}})),U={};function $(e,t=null){if(U[e])return void("function"==typeof t&&(U[e].push(t),t()));U[e]=[],"function"==typeof t&&U[e].push(t);const o=document.createElement("script");o.type="text/javascript",o.src=e,o.onload=()=>{U[e].forEach((e=>e()))},document.head.appendChild(o)}function N(t,o="myConvert"){const r={moduleCache:{vue:e},getFile:()=>F(t,{method:"GET"}).then((e=>e)),addStyle(e){const t=document.createElement("style");t.setAttribute("id",o),t.textContent=e;const r=document.head.getElementsByTagName("style")[document.head.getElementsByTagName("style").length-1]||null;document.head.insertBefore(t,r)}},n=_((()=>L(`${o}.vue`,r)));return g(n)}const z=k(),M={default:B("views/main/styleDefault.vue")},q=t("router",(()=>{const e=f(0),t=f([]);return{initialized:e,loadServerRouter:async(o=!1)=>{const r=C();if(o||0===e.value){e.value=1;const o=await F({url:"/sys/auth/userPermission",method:"GET"});t.value=o.data.menu[0].children,r.setUserData(o.data.user),V(t.value)}let n=r.userInfo();var s;return H.addRoute({path:"/"+z,name:z,meta:{icon:"add"},component:(s=n.user_config.theme?n.user_config.theme:"default",M[s]||(M[s]=N(s)),M[s]),children:t.value}),t.value}}})),G=k(),H=c({history:p("./"),routes:[{path:"/",redirect:`/${G}/welcome`},{path:"/login",name:"login",component:()=>S((()=>import("./XloginCqK-bZSU.js")),__vite__mapDeps([2,3,4,5,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]),import.meta.url)},{path:"/"+G,name:G,component:()=>S((()=>import("./XstyleDefaultB4IU6vd8.js")),__vite__mapDeps([23,3,4,9,1,6,7,8,10,11,12,13,14,15,16,17,18,19,5,20,21]),import.meta.url)},{path:"/:catchAll(.*)",meta:{closeTab:!0},component:()=>S((()=>import("./X404DLzpbq3O.js")),__vite__mapDeps([0,1]),import.meta.url)}]});H.beforeEach((async(e,t)=>{if(d.start(),e.meta.title&&(window.document.title=e.meta.title),"/login"===e.path)return!0;const o=q(),r=C();return o.initialized?(r.isLogIn()||r.logOut(),!0):(await o.loadServerRouter(!0),{path:e.path})})),H.afterEach((()=>{d.done()})),H.onError((()=>{d.done()}));const F=u.create({baseURL:W(),timeout:3e4});let Q,J=0;const K=()=>{J--,J<=0&&(clearTimeout(Q),X.emit("closeLoading"))};function W(){let e=[];return window.adminBox.Host&&e.push(window.adminBox.Host),window.adminBox.BackendRouterPrefix&&e.push(window.adminBox.BackendRouterPrefix),e.join("/")}F.interceptors.request.use((e=>{e.donNotShowLoading||(J++,Q&&clearTimeout(Q),Q=setTimeout((()=>{J>0&&X.emit("showLoading")}),400));const t=C();return e.headers={"Content-Type":"application/json; charset=utf-8",Authorization:"Bearer "+t.token(),"X-User-Type":t.userType(),"X-User-Id":t.userId(),...e.headers},e}),(e=>(e.config.donNotShowLoading||K(),r({showClose:!0,message:e,type:"error"}),e))),F.interceptors.response.use((e=>e.data),(e=>{if(e.config.donNotShowLoading||K(),e.response){switch(e.response.status){case 401:r({showClose:!0,message:e.response.data.msg,type:"error"}),H.push({name:"login",replace:!0});break;case 500:n.confirm(`\n        <p>检测到接口错误${e}</p>\n        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>\n        `,"接口报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"清理缓存",cancelButtonText:"取消"}).then((()=>{localStorage.clear(),H.push({name:"Login",replace:!0})}));break;case 404:n.confirm(`\n          <p>检测到接口错误${e}</p>\n          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>\n          `,"接口报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"我知道了",cancelButtonText:"取消"})}return e}n.confirm(`\n        <p>检测到请求错误</p>\n        <p>${e}</p>\n        `,"请求报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"稍后重试",cancelButtonText:"取消"})}));const Y={install:e=>{(async e=>{for(const r in l)e.component(r,l[r]);const t=await B("views/util/fileUpload.vue")(),o=await B("views/main/menuTree.vue")();e.component("FileUpload",t.default||t),e.component("MenuTree",o.default||o),e.config.globalProperties.$loadJS=$,e.config.globalProperties.$loadTMPL=N,e.config.globalProperties.$message=r,e.config.globalProperties.$http=F,e.config.globalProperties.$useUserStore=C,e.config.globalProperties.$useRouterStore=q})(e)}},Z={__name:"App",setup(e){let t="";if("en"===window.adminBox.Locale)t=a;else t=s;return(e,o)=>{const r=h("router-view"),n=h("el-config-provider");return v(),w(n,{locale:j(t)},{default:y((()=>[(v(),w(E,null,{default:y((()=>[T(r)])),_:1}))])),_:1},8,["locale"])}}};d.configure({showSpinner:!1,ease:"ease",speed:200});const ee=P(Z);ee.use(o()),ee.use(i),ee.use(Y),ee.use(H),ee.mount("#app");export{q as a,A as b,O as c,X as e,F as h,b as p,W as s,C as u,x as v};
