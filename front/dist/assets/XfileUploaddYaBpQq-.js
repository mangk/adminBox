import{K as S,u as se,M as C,r as i,c as n,d as c,e as _,g as p,Q as ue,f as o,w as u,h as D,F as M,A as F,k as L,l as $,s as B,R as ie,n as de,t as E,S as re}from"./XindexD6lHd35a.js";const ce=()=>S({url:"/sys/fileUpload/cfg",method:"GET"}),pe=(d=1,w=20,r={})=>S({url:"/sys/fileUpload/page",method:"POST",data:{page:d,page_size:w,query:r}}),ve=d=>S({url:"/sys/fileUpload",method:"DELETE",data:{id:d}}),fe={style:{display:"flex","flex-flow":"row wrap",padding:"var(--global-padding)"}},ge=["onMouseleave","onMouseover","onClick"],me={style:{"font-size":"30px","text-align":"center",margin:"var(--global-padding)","margin-bottom":"0"}},_e={style:{"font-size":"12px","line-height":"12px",padding:"3px","text-align":"center"}},he={style:{"font-size":"12px","line-height":"12px",padding:"3px",position:"absolute",left:"0",top:"0","background-color":"#409eff",color:"#fff","border-radius":"0 0 5px 0"}},ye={key:0},xe={class:"el-upload__text",style:{height:"40px","line-height":"40px","text-align":"center",overflow:"hidden",display:"flex","flex-flow":"row nowrap","justify-content":"center"}},Ce=p("em",null,"点击上传",-1),ze={__name:"fileUpload",props:{modelValue:{},modelModifiers:{},multiple:{default:1,type:Number},multipleModifiers:{},onlyPath:{default:!1,type:Boolean},onlyPathModifiers:{},suffixList:{default:[],type:Array},suffixListModifiers:{}},emits:["update:modelValue","update:multiple","update:onlyPath","update:suffixList"],setup(d){const w=se(),r=C(d,"modelValue"),h=C(d,"multiple"),O=C(d,"onlyPath"),N=C(d,"suffixList"),l=i([]),z=i({}),b=i("default"),y=i(!1),V=i([]),v=i(1),f=i(25),P=i(0),g=()=>{pe(v.value,f.value,{tag:N.value}).then(e=>{V.value=e.data.list,v.value=e.data.page,f.value=e.data.page_size,P.value=e.data.total})},T=e=>{if(h.value==1){l.value.length&&l.value[0].id==e.id?l.value=[]:l.value=[e],r.value=l.value,x();return}for(let t=0;t<l.value.length;t++)if(l.value[t].id==e.id){l.value.splice(t,1),r.value=l.value,x();return}if(l.value.length>=h.value){re({type:"warning",message:`最多选择 ${h.value} 个`});return}l.value.push(e),r.value=l.value,x()},j=e=>{for(const t in l.value)if(Object.hasOwnProperty.call(l.value,t)&&l.value[t].id==e)return"img-item-select"},A=(e,t,s)=>{g()},k=i(0),I=e=>{k.value=0},q=e=>{k.value=e},G=e=>{ve(e).then(t=>{for(let s=0;s<l.value.length;s++)l.value[s].id==e&&(l.value.splice(s,1),r.value=l.value);g()})},H=e=>{f.value=e,g()},K=e=>{v.value=e,g()},Q=()=>{ce().then(e=>{z.value=e.data}),g(),y.value=!0},R=()=>{x(),y.value=!1},x=()=>{if(l.value.length&&O.value)if(h.value==1)r.value=l.value[0].url;else{let e=[];for(let t=0;t<l.value.length;t++){const s=l.value[t];e.push(s.url)}r.value=e}};return(e,t)=>{const s=n("UploadFilled"),m=n("el-icon"),J=n("el-button"),W=n("el-pagination"),X=n("CircleClose"),Y=n("Document"),Z=n("el-image"),ee=n("el-radio-button"),le=n("el-radio-group"),te=n("upload-filled"),ae=n("el-upload"),oe=n("el-drawer");return c(),_(M,null,[p("div",{class:"upload-file-slot",onClick:Q},[ue(e.$slots,"default",{},()=>[o(J,{type:"primary"},{default:u(()=>[o(m,null,{default:u(()=>[o(s)]),_:1}),D(" 文件上传 ")]),_:1})])]),o(oe,{modelValue:y.value,"onUpdate:modelValue":t[3]||(t[3]=a=>y.value=a),direction:"rtl","before-close":R,"show-close":!1,size:"50%;",style:{"max-width":"600px"}},{header:u(()=>[o(W,{"current-page":v.value,"onUpdate:currentPage":t[0]||(t[0]=a=>v.value=a),"page-size":f.value,"onUpdate:pageSize":t[1]||(t[1]=a=>f.value=a),"page-sizes":[20,50,100,200],size:"small",layout:"total, prev, pager, next",total:P.value,onSizeChange:H,onCurrentChange:K},null,8,["current-page","page-size","total"])]),footer:u(()=>[Object.keys(z.value).length>1?(c(),_("div",ye,[o(le,{modelValue:b.value,"onUpdate:modelValue":t[2]||(t[2]=a=>b.value=a),size:"small"},{default:u(()=>[(c(!0),_(M,null,F(z.value,(a,ne)=>(c(),L(ee,{label:a.name,value:a.driver,key:a.driver},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])])):$("",!0),o(ae,{drag:"",action:`${B(ie)()}/sys/fileUpload/upload?driver=${b.value}`,headers:{Authorization:"Bearer "+B(w).token()},"show-file-list":!1,multiple:"","on-success":A},{default:u(()=>[p("div",xe,[o(m,{class:"el-icon--upload",style:{"font-size":"35px",margin:"0 var(--global-padding) 0 0"}},{default:u(()=>[o(te)]),_:1}),D(" 将文件拖到此处或"),Ce])]),_:1},8,["action","headers"])]),default:u(()=>[p("div",fe,[(c(!0),_(M,null,F(V.value,(a,ne)=>(c(),_("div",{class:de(["img-item",j(a.id)]),onMouseleave:U=>I(a.id),onMouseover:U=>q(a.id),onClick:U=>T(a)},[k.value==a.id?(c(),L(m,{key:0,class:"img-del",onClick:U=>G(a.id)},{default:u(()=>[o(X)]),_:2},1032,["onClick"])):$("",!0),o(Z,{style:{width:"100%"},src:a.url,fit:"scale-down",lazy:""},{error:u(()=>[p("div",me,[o(m,{style:{"font-size":"28px"}},{default:u(()=>[o(Y)]),_:1})]),p("div",_e,E(a.name.split(".")[0]),1),p("div",he,E(a.name.split(".")[1]),1)]),_:2},1032,["src"])],42,ge))),256))])]),_:1},8,["modelValue"])],64)}}};export{ze as default};
