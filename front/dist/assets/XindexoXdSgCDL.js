const __vite__fileDeps=["assets/X404DLzpbq3O.js","assets/X_plugin-vue_export-helperBCo6x5W8.js","assets/XloginB1PPt_Sm.js","assets/XvueCyDLjQZ5.js","assets/X@vueSdcp2x8a.js","assets/Xvue-routerCZMibURm.js","assets/XpiniaDkaaaSbL.js","assets/Xelement-plusViPJHclL.js","assets/Xlodash-esBg5u8xPa.js","assets/X@vueuseCSz_Rkca.js","assets/X@element-plusCAKmCgXq.js","assets/X@popperjsD3lHDW-0.js","assets/X@ctrlD2oWfImC.js","assets/Xdayjst441MSVg.js","assets/Xasync-validatorCuo4gI4y.js","assets/Xmemoize-oneDs0C_khL.js","assets/Xnormalize-wheel-esVn5vHDCm.js","assets/X@floating-uipMauM7H8.js","assets/XaxiosCURSphCx.js","assets/XmittCNZ6avp8.js","assets/XnprogressDgs0sPf-.js","assets/Xvue3-sfc-loaderDa78GGQm.js","assets/XmenuTreeDxIhcR46.js","assets/XstyleDefaultDeCGUGqB.js","assets/XstyleWhiteD_Sr7sUM.js","assets/XtabqfGuYUaV.js","assets/XapiC0QVwXxQ.js","assets/XmenuDPEX1v3Q.js","assets/XdepartmentDslPtuh_.js","assets/XdepartmentDNYrrHYj.js","assets/XpremissionSetlAVw4vjN.js","assets/XindexDdrtUNH3.js","assets/XmenuDWguuAQ5.js","assets/XroleCldJaBi4.js","assets/XroleDoI3DXQI.js","assets/XusererZSL9aV.js","assets/XfileUploadnTwkQ6ES.js","assets/XserverComponentBde-apHT.js","assets/XwelcomeCoBwhLTE.js"],__vite__mapDeps=i=>i.map(i=>__vite__fileDeps[i]);
import{V as e}from"./XvueCyDLjQZ5.js";import{d as t,c as o}from"./XpiniaDkaaaSbL.js";import{E as n,a as r,z as s,b as i,c as a}from"./Xelement-plusViPJHclL.js";import{Q as l}from"./X@element-plusCAKmCgXq.js";import{a as c}from"./XaxiosCURSphCx.js";import{m as u}from"./XmittCNZ6avp8.js";import{c as m,a as d}from"./Xvue-routerCZMibURm.js";import{N as p}from"./XnprogressDgs0sPf-.js";import{r as _,am as f,aY as g,ah as h,o as v,P as w,Q as y,aL as E,X as j,u as T,as as X}from"./X@vueSdcp2x8a.js";import{i as L}from"./Xvue3-sfc-loaderDa78GGQm.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCSz_Rkca.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver((e=>{for(const o of e)if("childList"===o.type)for(const e of o.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&t(e)})).observe(document,{childList:!0,subtree:!0})}function t(e){if(e.ep)return;e.ep=!0;const t=function(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerPolicy&&(t.referrerPolicy=e.referrerPolicy),"use-credentials"===e.crossOrigin?t.credentials="include":"anonymous"===e.crossOrigin?t.credentials="omit":t.credentials="same-origin",t}(e);fetch(e.href,t)}}();const P=u(),I={},S=function(e,t,o){let n=Promise.resolve();if(t&&t.length>0){document.getElementsByTagName("link");const e=document.querySelector("meta[property=csp-nonce]"),o=(null==e?void 0:e.nonce)||(null==e?void 0:e.getAttribute("nonce"));n=Promise.all(t.map((e=>{if((e=function(e){return"/"+e}(e))in I)return;I[e]=!0;const t=e.endsWith(".css"),n=t?'[rel="stylesheet"]':"";if(document.querySelector(`link[href="${e}"]${n}`))return;const r=document.createElement("link");return r.rel=t?"stylesheet":"modulepreload",t||(r.as="script",r.crossOrigin=""),r.href=e,o&&r.setAttribute("nonce",o),document.head.appendChild(r),t?new Promise(((t,o)=>{r.addEventListener("load",t),r.addEventListener("error",(()=>o(new Error(`Unable to preload CSS for ${e}`))))})):void 0})))}return n.then((()=>e())).catch((e=>{const t=new Event("vite:preloadError",{cancelable:!0});if(t.payload=e,window.dispatchEvent(t),!t.defaultPrevented)throw e}))},A=()=>q({url:"/sys/auth/permissionAll",method:"GET"}),O=(e,t)=>q({url:"/sys/auth/permissionGetByIdAndModule",method:"POST",data:{id:e,module:t}}),b=e=>q({url:"/sys/auth/permissionSave",method:"PUT",data:e}),x=()=>q({url:"/sys/verificationCode",method:"GET"}),R=Object.assign({"../views/404.vue":()=>S((()=>import("./X404DLzpbq3O.js")),__vite__mapDeps([0,1])),"../views/login.vue":()=>S((()=>import("./XloginB1PPt_Sm.js")),__vite__mapDeps([2,3,4,5,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21])),"../views/main/menuTree.vue":()=>S((()=>import("./XmenuTreeDxIhcR46.js")),__vite__mapDeps([22,3,4])),"../views/main/styleDefault.vue":()=>S((()=>import("./XstyleDefaultDeCGUGqB.js")),__vite__mapDeps([23,3,4,22,9,1,6,7,8,10,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/main/styleWhite.vue":()=>S((()=>import("./XstyleWhiteD_Sr7sUM.js")),__vite__mapDeps([24,3,4,22,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/main/tab.vue":()=>S((()=>import("./XtabqfGuYUaV.js")),__vite__mapDeps([25,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21])),"../views/setting/api.vue":()=>S((()=>import("./XapiC0QVwXxQ.js")),__vite__mapDeps([26,3,4,27,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/setting/department.vue":()=>S((()=>import("./XdepartmentDslPtuh_.js")),__vite__mapDeps([28,3,4,10,29,30,1,6,7,8,9,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/setting/index.vue":()=>S((()=>import("./XindexDdrtUNH3.js")),__vite__mapDeps([31,1])),"../views/setting/menu.vue":()=>S((()=>import("./XmenuDWguuAQ5.js")),__vite__mapDeps([32,3,4,27,10,1,6,7,8,9,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/setting/premissionSet.vue":()=>S((()=>import("./XpremissionSetlAVw4vjN.js")),__vite__mapDeps([30,3,4,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/setting/role.vue":()=>S((()=>import("./XroleCldJaBi4.js")),__vite__mapDeps([33,3,4,34,30,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21])),"../views/setting/user.vue":()=>S((()=>import("./XusererZSL9aV.js")),__vite__mapDeps([35,3,4,30,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21,29,34,36])),"../views/util/fileUpload.vue":()=>S((()=>import("./XfileUploadnTwkQ6ES.js")),__vite__mapDeps([36,3,4,7,8,9,10,11,12,13,14,15,16,17,6,18,19,5,20,21])),"../views/util/serverComponent.vue":()=>S((()=>import("./XserverComponentBde-apHT.js")),__vite__mapDeps([37,3,4,5,1])),"../views/welcome.vue":()=>S((()=>import("./XwelcomeCoBwhLTE.js")),__vite__mapDeps([38,3,4,6,7,8,9,10,11,12,13,14,15,16,17,18,19,5,20,21]))}),D=Object.assign({});function B(e){const t=Object.keys(R).filter((t=>t.replace("../","")===e))[0];return R[t]}function V(e){e.length&&e.forEach((e=>{e.component&&("views"===e.component.split("/")[0]?e.component=B(e.component):"plugins"===e.component.split("/")[0]&&(e.component=function(e){const t=Object.keys(D).filter((t=>t.replace("../","")===e))[0];return D[t]}(e.component))),e.children&&e.children.length&&(delete e.component,V(e.children))}))}function k(){var e="admin";return window._adminBox&&window._adminBox.backgroundPrefix&&(e=window._adminBox.backgroundPrefix),e}const C=t("user",(()=>{const e=_(0),t=_({}),o="x-token",n="x-user-type",r="x-user-id";return{initialized:e,userInfo:(o=!1)=>((o||0===e.value)&&(e.value=1),t.value),setUserData:e=>{t.value=e},token:()=>{var e=localStorage.getItem(o);if(e)return e;z.replace({name:"login"})},userType:()=>localStorage.getItem(n),userId:()=>localStorage.getItem(r),logIn:async s=>{const i=k();var a,l=await(a=s,q({url:"/sys/login",method:"POST",data:a}));if(0!=l.code)return"";t.value=l.data,e.value=1,await localStorage.setItem(o,l.data.jwt_token),await localStorage.setItem(n,"default"),await localStorage.setItem(r,l.data.id);const c=$();return await c.loadServerRouter(!0),l.data.user_config.home_page?l.data.user_config.home_page:`/${i}/welcome`},logOut:()=>{localStorage.clear(),window.location.href=window.location.origin},isLogIn:()=>localStorage.getItem(o)&&localStorage.getItem(n)&&localStorage.getItem(r)}})),U=k(),$=t("router",(()=>{const e=_(0),t=_([]);return{initialized:e,loadServerRouter:async(o=!1)=>{const n=C();if(o||0===e.value){e.value=1;const o=await q({url:"/sys/auth/userPermission",method:"GET"});t.value=o.data.menu[0].children,n.setUserData(o.data.user),V(t.value)}let r=n.userInfo(),s={default:B("views/main/styleDefault.vue"),white:B("views/main/styleWhite.vue")},i=s[r.user_config.theme?r.user_config.theme:"default"];return z.addRoute({path:"/"+U,name:U,meta:{icon:"add"},component:i||s.default,children:t.value}),t.value}}})),N=k(),z=m({history:d("/"),routes:[{path:"/",redirect:`/${N}/welcome`},{path:"/login",name:"login",component:()=>S((()=>import("./XloginB1PPt_Sm.js")),__vite__mapDeps([2,3,4,5,1,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21]))},{path:"/"+N,name:N,component:()=>S((()=>import("./XstyleDefaultDeCGUGqB.js")),__vite__mapDeps([23,3,4,22,9,1,6,7,8,10,11,12,13,14,15,16,17,18,19,5,20,21]))},{path:"/:catchAll(.*)",meta:{closeTab:!0},component:()=>S((()=>import("./X404DLzpbq3O.js")),__vite__mapDeps([0,1]))}]});z.beforeEach((async(e,t)=>{if(p.start(),e.meta.title&&(window.document.title=e.meta.title),"/login"===e.path)return!0;const o=$(),n=C();return o.initialized?(n.isLogIn()||n.logOut(),!0):(await o.loadServerRouter(!0),{path:e.path})})),z.afterEach((()=>{p.done()})),z.onError((()=>{p.done()}));const q=c.create({baseURL:W(),timeout:3e4});let M,G=0;const H=()=>{G--,G<=0&&(clearTimeout(M),P.emit("closeLoading"))};function W(){let e=[];return window.adminBox.Host&&e.push(window.adminBox.Host),window.adminBox.BackendRouterPrefix&&e.push(window.adminBox.BackendRouterPrefix),e.join("/")}function F(e,t=null){if(document.querySelector(`script[src="${e}"]`))return void("function"==typeof t&&t());const o=document.createElement("script");o.type="text/javascript",o.src=e,o.onload=()=>{"function"==typeof t&&t()},document.head.appendChild(o)}function Q(t,o="myConvert"){const n={moduleCache:{vue:e},getFile:()=>q(t,{method:"GET"}).then((e=>e)),addStyle(e){const t=document.createElement("style");t.setAttribute("id",o),t.textContent=e;const n=document.head.getElementsByTagName("style")[0]||null;document.head.insertBefore(t,n)}},r=g((()=>L(`${o}.vue`,n)));return f(r)}q.interceptors.request.use((e=>{e.donNotShowLoading||(G++,M&&clearTimeout(M),M=setTimeout((()=>{G>0&&P.emit("showLoading")}),400));const t=C();return e.headers={"Content-Type":"application/json; charset=utf-8",Authorization:"Bearer "+t.token(),"X-User-Type":t.userType(),"X-User-Id":t.userId(),...e.headers},e}),(e=>(e.config.donNotShowLoading||H(),n({showClose:!0,message:e,type:"error"}),e))),q.interceptors.response.use((e=>e.data),(e=>{if(e.config.donNotShowLoading||H(),e.response){switch(e.response.status){case 401:n({showClose:!0,message:e.response.data.msg,type:"error"}),z.push({name:"login",replace:!0});break;case 500:r.confirm(`\n        <p>检测到接口错误${e}</p>\n        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>\n        `,"接口报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"清理缓存",cancelButtonText:"取消"}).then((()=>{localStorage.clear(),z.push({name:"Login",replace:!0})}));break;case 404:r.confirm(`\n          <p>检测到接口错误${e}</p>\n          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>\n          `,"接口报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"我知道了",cancelButtonText:"取消"})}return e}r.confirm(`\n        <p>检测到请求错误</p>\n        <p>${e}</p>\n        `,"请求报错",{dangerouslyUseHTMLString:!0,distinguishCancelAndClose:!0,confirmButtonText:"稍后重试",cancelButtonText:"取消"})}));const J={install:e=>{(async e=>{for(const o in l)e.component(o,l[o]);const t=await B("views/util/fileUpload.vue")();e.component("FileUpload",t.default||t),e.config.globalProperties.$loadJS=F,e.config.globalProperties.$loadTMPL=Q,e.config.globalProperties.$message=n,e.config.globalProperties.$http=q})(e)}},K={__name:"App",setup(e){let t="";if("en"===window.adminBox.Locale)t=i;else t=s;return(e,o)=>{const n=h("router-view"),r=h("el-config-provider");return v(),w(r,{locale:T(t)},{default:y((()=>[(v(),w(E,null,{default:y((()=>[j(n)])),_:1}))])),_:1},8,["locale"])}}};p.configure({showSpinner:!1,ease:"ease",speed:200});const Y=X(K);Y.use(o()),Y.use(a),Y.use(J),Y.use(z),Y.mount("#app");export{$ as a,O as b,b as c,P as e,q as h,A as p,W as s,C as u,x as v};