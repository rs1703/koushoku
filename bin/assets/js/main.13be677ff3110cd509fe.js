!function(){"use strict";function t(){var n,e,o,t,r,c,i,a,d,u,s,l,f=document.getElementById("reader");f&&(t=window.location.pathname.split("/"),p=t[2],g=t[3],n=Number(document.querySelector(".total").textContent),e=document.querySelectorAll(".current"),o=Number(e[0].textContent),t=document.querySelectorAll(".first"),r=document.querySelectorAll(".last"),c=document.querySelectorAll(".prev"),i=document.querySelectorAll(".next"),a=f.querySelector(".page a"),d=a.querySelector("img"),u={current:!1},s=a.getBoundingClientRect(),m.push.apply(m,Array.from({length:n},function(t,e){return{isViewing:e+1===o}})),l=function(t){o=t,a.scrollIntoView({block:"start",inline:"start"}),window.requestAnimationFrame(function(){m.find(function(t){return t.isViewing}).isViewing=!1,m[o-1].isViewing=!0,a.href="/archive/".concat(p,"/").concat(g,"/").concat(o),d.src="/data/".concat(p,"/").concat(o,".jpg"),e.forEach(function(t){return t.textContent=o.toString()}),window.history.replaceState(null,"","/archive/".concat(p,"/").concat(g,"/").concat(o))}),c.forEach(function(t){t.href="/archive/".concat(p,"/").concat(g,"/").concat(o-1)}),i.forEach(function(t){t.href="/archive/".concat(p,"/").concat(g,"/").concat(o+1)})},t.forEach(function(t){t.addEventListener("click",function(t){t.preventDefault(),t.stopPropagation(),t.stopImmediatePropagation(),1!==o&&l(1)})}),r.forEach(function(t){t.addEventListener("click",function(t){t.preventDefault(),t.stopPropagation(),t.stopImmediatePropagation(),o!==n&&l(n)})}),c.forEach(function(t){t.addEventListener("click",function(t){t.preventDefault(),t.stopPropagation(),t.stopImmediatePropagation(),1!==o&&l(o-1)})}),i.forEach(function(t){t.addEventListener("click",function(t){t.preventDefault(),t.stopPropagation(),t.stopImmediatePropagation(),o!==n&&l(o+1)})}),a.addEventListener("click",function(t){var e;t.preventDefault(),t.stopPropagation(),t.stopImmediatePropagation(),u.current||(u.current=!0,e=t.target,t.screenX-s.x<=e.clientWidth/2?1!==o&&l(o-1):o!==n&&l(o+1),u.current=!1)}),new MutationObserver(function(){return v()}).observe(a,{childList:!0,attributes:!0}),a.scrollIntoView({block:"start",inline:"start"}),v())}function e(){"serviceWorker"in navigator&&navigator.serviceWorker.register("/serviceWorker.js"),t()}var p,g,m=[],v=function(){for(var r=m.findIndex(function(t){return t.isViewing}),t=function(t){var e=m[t];if(e.isPreloading||e.isPreloaded||t<Math.max(0,r-3)||t>Math.max(0,r+3))return"continue";e.isPreloading=!0;function n(){e.isPreloading=!1,e.isPreloaded=!0}var o=document.createElement("img");o.src="/data/".concat(p,"/").concat(t+1,".jpg");o.complete?n():(o.addEventListener("load",n),o.addEventListener("error",function(){return e.isPreloading=!1}))},e=0;e<m.length;e++)t(e)};"complete"===document.readyState?e():document.addEventListener("DOMContentLoaded",e)}();