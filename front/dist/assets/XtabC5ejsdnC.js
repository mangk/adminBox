import{z as B,b as M,r,A as b,B as p,C as R,f as O,o as S,c as _,d as L,w as N,F as D,y as F,j as T,e as d,t as z,D as H,E as P,G as U,H as G}from"./XindexrDa97Uv6.js";const K={class:"router-history"},X=["tab"],W={__name:"tab",setup(Y){const m=B(),i=M(),n=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),s=r([]),u=r(""),c=r(!1);r({});const g=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),I=r(0),w=r(0),x=r(!1),v=r(""),f=r("/admin/welcome"),J=e=>{if(s.value.length===1&&m.name===f.value)return!1;let a="";e.srcElement.nodeName==="SPAN"?a=e.srcElement.offsetParent.id:a=e.srcElement.id,a&&(c.value=!0,I.value=e.clientX,w.value=e.clientY+10,v.value=a.substring(4))},y=()=>{s.value=[],i.push({path:f.value}),c.value=!1,sessionStorage.setItem("historys",JSON.stringify(s.value))},V=()=>{let e;const a=s.value.findIndex(l=>(n(l)===v.value&&(e=l),n(l)===v.value)),t=s.value.findIndex(l=>n(l)===u.value);s.value.splice(0,a),a>t&&i.push(e),sessionStorage.setItem("historys",JSON.stringify(s.value))},A=()=>{let e;const a=s.value.findIndex(l=>(n(l)===v.value&&(e=l),n(l)===v.value)),t=s.value.findIndex(l=>n(l)===u.value);s.value.splice(a+1,s.value.length),a<t&&i.push(e),sessionStorage.setItem("historys",JSON.stringify(s.value))},C=()=>{let e;s.value=s.value.filter(a=>(n(a)===v.value&&(e=a),n(a)===v.value)),i.push(e),sessionStorage.setItem("historys",JSON.stringify(s.value))},E=(e,a)=>{if(e.name!==a.name||Object.keys(e.query).length!==Object.keys(a.query).length||Object.keys(e.params).length!==Object.keys(a.params).length)return!1;for(const t in e.query)if(e.query[t]!==a.query[t])return!1;for(const t in e.params)if(e.params[t]!==a.params[t])return!1;return!0},k=e=>{if(!s.value.some(a=>E(a,e))){const a={};a.name=e.name,a.meta={...e.meta},delete a.meta.matched,a.query=e.query,a.params=e.params,s.value.push(a)}window.sessionStorage.setItem("activeValue",n(e))},h=r({}),j=e=>{var l;const a=(l=e==null?void 0:e.props)==null?void 0:l.name;if(!a)return;const t=h.value[a];i.push({name:t.name,query:t.query,params:t.params})},q=e=>{const a=s.value.findIndex(t=>n(t)===e);n(m)===e&&(s.value.length===1?i.push({path:f.value}):a<s.value.length-1?i.push({name:s.value[a+1].name,query:s.value[a+1].query,params:s.value[a+1].params}):i.push({name:s.value[a-1].name,query:s.value[a-1].query,params:s.value[a-1].params})),s.value.splice(a,1),s.value.length===0&&i.push({path:f.value})};return b(()=>c.value,()=>{c.value?document.body.addEventListener("click",()=>{c.value=!1}):document.body.removeEventListener("click",()=>{c.value=!1})}),b(()=>m,(e,a)=>{e.name!=="login"&&(s.value=s.value.filter(t=>!t.meta.closeTab),k(e),sessionStorage.setItem("historys",JSON.stringify(s.value)),u.value=window.sessionStorage.getItem("activeValue"))},{deep:!0}),b(()=>s.value,()=>{sessionStorage.setItem("historys",JSON.stringify(s.value)),h.value={},s.value.forEach(e=>{h.value[n(e)]=e}),p.emit("setKeepAlive",s.value)},{deep:!0}),(()=>{p.on("closeThisPage",()=>{q(g(m))}),p.on("closeAllPage",()=>{y()}),p.on("mobile",a=>{x.value=a});const e=[{name:f.value,meta:{title:"欢迎"},query:{},params:{}}];s.value=JSON.parse(sessionStorage.getItem("historys"))||e,window.sessionStorage.getItem("activeValue")?u.value=window.sessionStorage.getItem("activeValue"):u.value=n(m),k(m),window.sessionStorage.getItem("needCloseAll")==="true"&&(y(),window.sessionStorage.removeItem("needCloseAll"))})(),R(()=>{p.off("collapse"),p.off("mobile")}),(e,a)=>{const t=O("el-tab-pane"),l=O("el-tabs");return S(),_("div",K,[L(l,{modelValue:u.value,"onUpdate:modelValue":a[0]||(a[0]=o=>u.value=o),closable:!(s.value.length===1&&e.$route.name===f.value),type:"card",onContextmenu:a[1]||(a[1]=H(o=>J(o),["prevent"])),onTabClick:j,onTabRemove:q,class:"app-tab"},{default:N(()=>[(S(!0),_(D,null,F(s.value,o=>(S(),T(t,{key:g(o),label:o.meta.title,name:g(o),tab:o},{label:N(()=>[d("span",{tab:o},z(o.meta.title),9,X)]),_:2},1032,["label","name","tab"]))),128))]),_:1},8,["modelValue","closable"]),P(d("ul",{style:G({left:I.value+"px",top:w.value+"px"}),class:"contextmenu"},[d("li",{onClick:y},"关闭所有"),d("li",{onClick:V},"关闭左侧"),d("li",{onClick:A},"关闭右侧"),d("li",{onClick:C},"关闭其他")],4),[[U,c.value]])])}}};export{W as default};
