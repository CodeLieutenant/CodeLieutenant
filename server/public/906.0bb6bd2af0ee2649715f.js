(self.webpackChunktemplate=self.webpackChunktemplate||[]).push([[906],{2956:(e,t,r)=>{"use strict";r.d(t,{d:()=>a});var n=function(){return(n=Object.assign||function(e){for(var t,r=1,n=arguments.length;r<n;r++)for(var a in t=arguments[r])Object.prototype.hasOwnProperty.call(t,a)&&(e[a]=t[a]);return e}).apply(this,arguments)};function a(e,t,r,a){return o=this,i=void 0,c=function(){var o,i;return function(e,t){var r,n,a,o,i={label:0,sent:function(){if(1&a[0])throw a[1];return a[1]},trys:[],ops:[]};return o={next:s(0),throw:s(1),return:s(2)},"function"==typeof Symbol&&(o[Symbol.iterator]=function(){return this}),o;function s(o){return function(s){return function(o){if(r)throw new TypeError("Generator is already executing.");for(;i;)try{if(r=1,n&&(a=2&o[0]?n.return:o[0]?n.throw||((a=n.return)&&a.call(n),0):n.next)&&!(a=a.call(n,o[1])).done)return a;switch(n=0,a&&(o=[2&o[0],a.value]),o[0]){case 0:case 1:a=o;break;case 4:return i.label++,{value:o[1],done:!1};case 5:i.label++,n=o[1],o=[0];continue;case 7:o=i.ops.pop(),i.trys.pop();continue;default:if(!((a=(a=i.trys).length>0&&a[a.length-1])||6!==o[0]&&2!==o[0])){i=0;continue}if(3===o[0]&&(!a||o[1]>a[0]&&o[1]<a[3])){i.label=o[1];break}if(6===o[0]&&i.label<a[1]){i.label=a[1],a=o;break}if(a&&i.label<a[2]){i.label=a[2],i.ops.push(o);break}a[2]&&i.ops.pop(),i.trys.pop();continue}o=t.call(e,i)}catch(e){o=[6,e],n=0}finally{r=a=0}if(5&o[0])throw o[1];return{value:o[0]?o[1]:void 0,done:!0}}([o,s])}}}(this,(function(s){switch(s.label){case 0:if(a?(a.method=t,a.headers=n({"Content-Type":"application/json",Accept:"application/json","X-Requested-With":"XMLHttpRequest"},a.headers)):a={method:t,headers:{"Content-Type":"application/json",Accept:"application/json","X-Requested-With":"XMLHttpRequest"}},"Content-Type"in a.headers)switch(a.headers["Content-Type"]){case"application/json":a.body=JSON.stringify(r);break;case"multipart/form-data":for(i in o=new FormData,r)o.append(i,r[i]);a.body=o}return[4,fetch(e,a)];case 1:return[2,s.sent()]}}))},new((s=void 0)||(s=Promise))((function(e,t){function r(e){try{a(c.next(e))}catch(e){t(e)}}function n(e){try{a(c.throw(e))}catch(e){t(e)}}function a(t){var a;t.done?e(t.value):(a=t.value,a instanceof s?a:new s((function(e){e(a)}))).then(r,n)}a((c=c.apply(o,i||[])).next())}));var o,i,s,c}},7906:(e,t,r)=>{"use strict";r.r(t),r.d(t,{subscribeFormHandler:()=>u});var n=r(9501),a=r(2956),o=function(e,t,r,n){return new(r||(r=Promise))((function(a,o){function i(e){try{c(n.next(e))}catch(e){o(e)}}function s(e){try{c(n.throw(e))}catch(e){o(e)}}function c(e){var t;e.done?a(e.value):(t=e.value,t instanceof r?t:new r((function(e){e(t)}))).then(i,s)}c((n=n.apply(e,t||[])).next())}))},i=function(e,t){var r,n,a,o,i={label:0,sent:function(){if(1&a[0])throw a[1];return a[1]},trys:[],ops:[]};return o={next:s(0),throw:s(1),return:s(2)},"function"==typeof Symbol&&(o[Symbol.iterator]=function(){return this}),o;function s(o){return function(s){return function(o){if(r)throw new TypeError("Generator is already executing.");for(;i;)try{if(r=1,n&&(a=2&o[0]?n.return:o[0]?n.throw||((a=n.return)&&a.call(n),0):n.next)&&!(a=a.call(n,o[1])).done)return a;switch(n=0,a&&(o=[2&o[0],a.value]),o[0]){case 0:case 1:a=o;break;case 4:return i.label++,{value:o[1],done:!1};case 5:i.label++,n=o[1],o=[0];continue;case 7:o=i.ops.pop(),i.trys.pop();continue;default:if(!((a=(a=i.trys).length>0&&a[a.length-1])||6!==o[0]&&2!==o[0])){i=0;continue}if(3===o[0]&&(!a||o[1]>a[0]&&o[1]<a[3])){i.label=o[1];break}if(6===o[0]&&i.label<a[1]){i.label=a[1],a=o;break}if(a&&i.label<a[2]){i.label=a[2],i.ops.push(o);break}a[2]&&i.ops.pop(),i.trys.pop();continue}o=t.call(e,i)}catch(e){o=[6,e],n=0}finally{r=a=0}if(5&o[0])throw o[1];return{value:o[0]?o[1]:void 0,done:!0}}([o,s])}}},s=r(6455),c=(0,n.Ry)().shape({name:(0,n.Z_)().required().max(50),email:(0,n.Z_)().required().email().max(150)}),u=function(e,t,r,u){return function(l){return o(void 0,void 0,void 0,(function(){var f,p,d,h,m;return i(this,(function(b){switch(b.label){case 0:return l.preventDefault(),f=document.getElementById(e),p=document.getElementById(t),d=document.getElementById(r),h=document.getElementById(u),[4,(y={name:f.value,email:p.value},o(void 0,void 0,void 0,(function(){var e,t,r;return i(this,(function(o){switch(o.label){case 0:return o.trys.push([0,4,,5]),[4,c.validate(y,{recursive:!0,abortEarly:!1})];case 1:return o.sent(),[4,(0,a.d)("/subscribe","POST",y)];case 2:return[4,o.sent().json()];case 3:return[2,{id:(e=o.sent()).id,name:e.name,email:e.email,createdAt:new Date(e.createdAt)}];case 4:return(t=o.sent())instanceof n.p8?(r={nameError:"",emailError:""},t.inner.forEach((function(e){"name"===e.path?r.nameError=e.errors[0]:"email"===e.path&&(r.emailError=e.errors[0])})),[2,r]):[2,{message:"Try again please"}];case 5:return[2]}}))})))];case 1:return"message"in(m=b.sent())?(s.fire({title:"Error",text:m.message,icon:"error",timerProgressBar:!0}),[2]):"nameError"in m&&"emailError"in m?(""!==m.nameError&&(f.classList.add("input-error"),d.classList.remove("hidden"),d.innerText=m.nameError),""!==m.emailError&&(p.classList.add("input-error"),h.classList.remove("hidden"),h.innerText=m.emailError),setTimeout((function(){d.classList.add("hidden"),h.classList.add("hidden"),p.classList.remove("input-error"),f.classList.remove("input-error")}),4e3),[2]):(gtag("event","subscribe",{event_category:"subscription",event_label:"New user subscribed to news letters"}),s.fire({title:"Success",text:"You have successfully subscribed to newsletters",icon:"success",timerProgressBar:!0}),[2])}var y}))}))}}}}]);