import{J as S,u as ie,N as w,r as i,f as n,o as c,c as _,e as p,R as de,d as o,w as u,g as D,F as P,z as L,j as $,k as F,q as re,n as ce,t as A,S as pe}from"./Xindexa4Gb5-_L.js";const E=window.adminX?window.adminX:{},V=E.BackendPrefix?"/"+E.BackendPrefix:"",fe=()=>S({url:V+"/sys/fileUpload/cfg",method:"GET"}),ve=(d=1,C=20,r={})=>S({url:V+"/sys/fileUpload/page",method:"POST",data:{page:d,page_size:C,query:r}}),me=d=>S({url:V+"/sys/fileUpload",method:"DELETE",data:{id:d}}),ge={style:{display:"flex","flex-flow":"row wrap",padding:"var(--global-padding)"}},_e=["onMouseleave","onMouseover","onClick"],he={style:{"font-size":"30px","text-align":"center",margin:"var(--global-padding)","margin-bottom":"0"}},xe={style:{"font-size":"12px","line-height":"12px",padding:"3px","text-align":"center"}},ye={style:{"font-size":"12px","line-height":"12px",padding:"3px",position:"absolute",left:"0",top:"0","background-color":"#409eff",color:"#fff","border-radius":"0 0 5px 0"}},we={key:0},Ce={class:"el-upload__text",style:{height:"40px","line-height":"40px","text-align":"center",overflow:"hidden",display:"flex","flex-flow":"row nowrap","justify-content":"center"}},ke=p("em",null,"点击上传",-1),Ue={__name:"fileUpload",props:{modelValue:{},modelModifiers:{},multiple:{default:1,type:Number},multipleModifiers:{},onlyPath:{default:!1,type:Boolean},onlyPathModifiers:{},suffixList:{default:[],type:Array},suffixListModifiers:{}},emits:["update:modelValue","update:multiple","update:onlyPath","update:suffixList"],setup(d){const C=ie(),r=w(d,"modelValue"),h=w(d,"multiple"),N=w(d,"onlyPath"),O=w(d,"suffixList"),l=i([]),k=i({}),z=i("default"),x=i(!1),y=window.adminX?window.adminX:{},j=(y.RunAt?y.RunAt:"")+(y.BackendPrefix?"/"+y.BackendPrefix:""),M=i([]),f=i(1),v=i(25),B=i(0),m=()=>{ve(f.value,v.value,{tag:O.value}).then(e=>{M.value=e.data.list,f.value=e.data.page,v.value=e.data.page_size,B.value=e.data.total})},T=e=>{if(h.value==1){l.value.length&&l.value[0].id==e.id?l.value=[]:l.value=[e],r.value=l.value;return}for(let t=0;t<l.value.length;t++)if(l.value[t].id==e.id){l.value.splice(t,1),r.value=l.value;return}if(l.value.length>=h.value){pe({type:"warning",message:`最多选择 ${h.value} 个`});return}l.value.push(e),r.value=l.value},X=e=>{for(const t in l.value)if(Object.hasOwnProperty.call(l.value,t)&&l.value[t].id==e)return"img-item-select"},R=(e,t,s)=>{m()},b=i(0),q=e=>{b.value=0},I=e=>{b.value=e},G=e=>{me(e).then(t=>{for(let s=0;s<l.value.length;s++)l.value[s].id==e&&(l.value.splice(s,1),r.value=l.value);m()})},H=e=>{v.value=e,m()},J=e=>{f.value=e,m()},K=()=>{fe().then(e=>{k.value=e.data}),m(),x.value=!0},Q=()=>{if(l.value.length&&N.value)if(h.value==1)r.value=l.value[0].url;else{let e=[];for(let t=0;t<l.value.length;t++){const s=l.value[t];e.push(s.url)}r.value=e}x.value=!1};return(e,t)=>{const s=n("UploadFilled"),g=n("el-icon"),W=n("el-button"),Y=n("el-pagination"),Z=n("CircleClose"),ee=n("Document"),le=n("el-image"),te=n("el-radio-button"),ae=n("el-radio-group"),oe=n("upload-filled"),ne=n("el-upload"),se=n("el-drawer");return c(),_(P,null,[p("div",{class:"upload-file-slot",onClick:K},[de(e.$slots,"default",{},()=>[o(W,{type:"primary"},{default:u(()=>[o(g,null,{default:u(()=>[o(s)]),_:1}),D(" 文件上传 ")]),_:1})])]),o(se,{modelValue:x.value,"onUpdate:modelValue":t[3]||(t[3]=a=>x.value=a),direction:"rtl","before-close":Q,"show-close":!1,size:"50%;",style:{"max-width":"600px"}},{header:u(()=>[o(Y,{"current-page":f.value,"onUpdate:currentPage":t[0]||(t[0]=a=>f.value=a),"page-size":v.value,"onUpdate:pageSize":t[1]||(t[1]=a=>v.value=a),"page-sizes":[20,50,100,200],size:"small",layout:"total, prev, pager, next",total:B.value,onSizeChange:H,onCurrentChange:J},null,8,["current-page","page-size","total"])]),footer:u(()=>[Object.keys(k.value).length>1?(c(),_("div",we,[o(ae,{modelValue:z.value,"onUpdate:modelValue":t[2]||(t[2]=a=>z.value=a),size:"small"},{default:u(()=>[(c(!0),_(P,null,L(k.value,(a,ue)=>(c(),$(te,{label:a.name,value:a.driver,key:a.driver},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])])):F("",!0),o(ne,{drag:"",action:`${j}/sys/fileUpload/upload?driver=${z.value}`,headers:{Authorization:"Bearer "+re(C).token()},"show-file-list":!1,multiple:"","on-success":R},{default:u(()=>[p("div",Ce,[o(g,{class:"el-icon--upload",style:{"font-size":"35px",margin:"0 var(--global-padding) 0 0"}},{default:u(()=>[o(oe)]),_:1}),D(" 将文件拖到此处或"),ke])]),_:1},8,["action","headers"])]),default:u(()=>[p("div",ge,[(c(!0),_(P,null,L(M.value,(a,ue)=>(c(),_("div",{class:ce(["img-item",X(a.id)]),onMouseleave:U=>q(a.id),onMouseover:U=>I(a.id),onClick:U=>T(a)},[b.value==a.id?(c(),$(g,{key:0,class:"img-del",onClick:U=>G(a.id)},{default:u(()=>[o(Z)]),_:2},1032,["onClick"])):F("",!0),o(le,{style:{width:"100%"},src:a.url,fit:"scale-down",lazy:""},{error:u(()=>[p("div",he,[o(g,{style:{"font-size":"28px"}},{default:u(()=>[o(ee)]),_:1})]),p("div",xe,A(a.name.split(".")[0]),1),p("div",ye,A(a.name.split(".")[1]),1)]),_:2},1032,["src"])],42,_e))),256))])]),_:1},8,["modelValue"])],64)}}};export{Ue as default};
