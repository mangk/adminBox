import"./XvueCyDLjQZ5.js";import{e}from"./XindexCDFV15DS.js";import{b as a,u as s}from"./Xvue-routerCZMibURm.js";import{r as t,f as l,E as n,ah as u,o,c as r,X as i,Q as m,F as v,a8 as p,P as c,a as d,W as y,a4 as g,R as f,Y as h,L as b}from"./X@vueSdcp2x8a.js";import"./XpiniaDkaaaSbL.js";import"./Xelement-plusJnkMASPo.js";import"./Xlodash-esBg5u8xPa.js";import"./X@vueuseCSz_Rkca.js";import"./X@element-plusCAKmCgXq.js";import"./X@popperjsD3lHDW-0.js";import"./X@ctrlD2oWfImC.js";import"./Xdayjst441MSVg.js";import"./Xasync-validatorCuo4gI4y.js";import"./Xmemoize-oneDs0C_khL.js";import"./Xnormalize-wheel-esVn5vHDCm.js";import"./X@floating-uipMauM7H8.js";import"./Xvue3-sfc-loaderDa78GGQm.js";import"./XaxiosCURSphCx.js";import"./XmittCNZ6avp8.js";import"./XnprogressDgs0sPf-.js";const j={class:"router-history"},S=["tab"],X={__name:"tab",setup(X){const I=a(),q=s(),O=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),w=t([]),k=t(""),N=t(!1);t({});const x=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),J=t(0),C=t(0),E=t(!1),V=t(""),A=t("/admin/welcome"),P=()=>{w.value=[],q.push({path:A.value}),N.value=!1,sessionStorage.setItem("historys",JSON.stringify(w.value))},T=()=>{let e;const a=w.value.findIndex((a=>(O(a)===V.value&&(e=a),O(a)===V.value))),s=w.value.findIndex((e=>O(e)===k.value));w.value.splice(0,a),a>s&&q.push(e),sessionStorage.setItem("historys",JSON.stringify(w.value))},_=()=>{let e;const a=w.value.findIndex((a=>(O(a)===V.value&&(e=a),O(a)===V.value))),s=w.value.findIndex((e=>O(e)===k.value));w.value.splice(a+1,w.value.length),a<s&&q.push(e),sessionStorage.setItem("historys",JSON.stringify(w.value))},L=()=>{let e;w.value=w.value.filter((a=>(O(a)===V.value&&(e=a),O(a)===V.value))),q.push(e),sessionStorage.setItem("historys",JSON.stringify(w.value))},z=e=>{if(!w.value.some((a=>((e,a)=>{if(e.name!==a.name)return!1;if(Object.keys(e.query).length!==Object.keys(a.query).length||Object.keys(e.params).length!==Object.keys(a.params).length)return!1;for(const s in e.query)if(e.query[s]!==a.query[s])return!1;for(const s in e.params)if(e.params[s]!==a.params[s])return!1;return!0})(a,e)))){const a={};a.name=e.name,a.meta={...e.meta},delete a.meta.matched,a.query=e.query,a.params=e.params,w.value.push(a)}window.sessionStorage.setItem("activeValue",O(e))},R=t({}),Y=e=>{var a;const s=null==(a=null==e?void 0:e.props)?void 0:a.name;if(!s)return;const t=R.value[s];q.push({name:t.name,query:t.query,params:t.params})},F=e=>{const a=w.value.findIndex((a=>O(a)===e));O(I)===e&&(1===w.value.length?q.push({path:A.value}):a<w.value.length-1?q.push({name:w.value[a+1].name,query:w.value[a+1].query,params:w.value[a+1].params}):q.push({name:w.value[a-1].name,query:w.value[a-1].query,params:w.value[a-1].params})),w.value.splice(a,1),0===w.value.length&&q.push({path:A.value})};l((()=>N.value),(()=>{N.value?document.body.addEventListener("click",(()=>{N.value=!1})):document.body.removeEventListener("click",(()=>{N.value=!1}))})),l((()=>I),((e,a)=>{"login"!==e.name&&(w.value=w.value.filter((e=>!e.meta.closeTab)),z(e),sessionStorage.setItem("historys",JSON.stringify(w.value)),k.value=window.sessionStorage.getItem("activeValue"))}),{deep:!0}),l((()=>w.value),(()=>{sessionStorage.setItem("historys",JSON.stringify(w.value)),R.value={},w.value.forEach((e=>{R.value[O(e)]=e})),e.emit("setKeepAlive",w.value)}),{deep:!0});return(()=>{e.on("closeThisPage",(()=>{F(x(I))})),e.on("closeAllPage",(()=>{P()})),e.on("mobile",(e=>{E.value=e}));const a=[{name:A.value,meta:{title:"欢迎"},query:{},params:{}}];w.value=JSON.parse(sessionStorage.getItem("historys"))||a,window.sessionStorage.getItem("activeValue")?k.value=window.sessionStorage.getItem("activeValue"):k.value=O(I),z(I),"true"===window.sessionStorage.getItem("needCloseAll")&&(P(),window.sessionStorage.removeItem("needCloseAll"))})(),n((()=>{e.off("collapse"),e.off("mobile")})),(e,a)=>{const s=u("el-tab-pane"),t=u("el-tabs");return o(),r("div",j,[i(t,{modelValue:k.value,"onUpdate:modelValue":a[0]||(a[0]=e=>k.value=e),closable:!(1===w.value.length&&e.$route.name===A.value),type:"card",onContextmenu:a[1]||(a[1]=g((e=>(e=>{if(1===w.value.length&&I.name===A.value)return!1;let a="";a="SPAN"===e.srcElement.nodeName?e.srcElement.offsetParent.id:e.srcElement.id,a&&(N.value=!0,J.value=e.clientX,C.value=e.clientY+10,V.value=a.substring(4))})(e)),["prevent"])),onTabClick:Y,onTabRemove:F,class:"app-tab"},{default:m((()=>[(o(!0),r(v,null,p(w.value,(e=>(o(),c(s,{key:x(e),label:e.meta.title,name:x(e),tab:e},{label:m((()=>[d("span",{tab:e},y(e.meta.title),9,S)])),_:2},1032,["label","name","tab"])))),128))])),_:1},8,["modelValue","closable"]),f(d("ul",{style:b({left:J.value+"px",top:C.value+"px"}),class:"contextmenu"},[d("li",{onClick:P},"关闭所有"),d("li",{onClick:T},"关闭左侧"),d("li",{onClick:_},"关闭右侧"),d("li",{onClick:L},"关闭其他")],4),[[h,N.value]])])}}};export{X as default};
