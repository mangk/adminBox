import{c as o,d as t,e as s,A as h,F as c,k as n,w as a,J as i,l,g as _,t as d,f as k}from"./XindexDO8YR9aY.js";const B={__name:"menuTree",props:["menus"],setup(m){return(x,y)=>{const r=o("el-icon"),u=o("MenuTree",!0),p=o("el-sub-menu"),f=o("el-menu-item");return t(!0),s(c,null,h(m.menus,e=>(t(),s(c,null,[e.children&&e.children.length&&!e.hidden?(t(),n(p,{index:e.name,key:e.id},{title:a(()=>[e.meta.icon?(t(),n(r,{key:0},{default:a(()=>[(t(),n(i(e.meta.icon)))]),_:2},1024)):l("",!0),_("span",null,d(e.meta.title),1)]),default:a(()=>[k(u,{menus:e.children},null,8,["menus"])]),_:2},1032,["index"])):e.hidden?l("",!0):(t(),n(f,{key:1,index:e.name,route:e},{default:a(()=>[e.meta.icon?(t(),n(r,{key:0},{default:a(()=>[(t(),n(i(e.meta.icon)))]),_:2},1024)):l("",!0),_("span",null,d(e.meta.title),1)]),_:2},1032,["index","route"]))],64))),256)}}};export{B as default};