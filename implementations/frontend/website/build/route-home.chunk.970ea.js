webpackJsonp([0],{UlEV:function(){},ZAL5:function(t){t.exports={home:"home__MseGd",cardHeader:"cardHeader__2Vd1U",cardBody:"cardBody__fYYFu"}},aqQ4:function(){},"iOg+":function(t,e,a){"use strict";function n(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!=typeof e&&"function"!=typeof e?t:e}function r(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+typeof e);t.prototype=Object.create(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(Object.setPrototypeOf?Object.setPrototypeOf(t,e):t.__proto__=e)}Object.defineProperty(e,"__esModule",{value:!0}),a.d(e,"default",function(){return v});var o=a("KM04"),u=(a.n(o),a("sJaT")),i=a.n(u),c=a("UlEV"),d=(a.n(c),a("aqQ4")),l=(a.n(d),a("ZAL5")),s=a.n(l),f=Object(o.h)("h1",null,"Home route"),p=Object(o.h)("h2",{class:" mdc-typography--title"},"Home card"),m=Object(o.h)("div",{class:" mdc-typography--caption"},"Welcome to home route"),h=Object(o.h)(i.a.Actions,null,Object(o.h)(i.a.ActionButton,null,"Turn On"),Object(o.h)(i.a.ActionButton,null,"Turn Off")),v=function(t){function e(){return n(this,t.apply(this,arguments))}return r(e,t),e.prototype.render=function(){return Object(o.h)("div",{class:s.a.home+" page"},f,Object(o.h)(i.a,null,Object(o.h)("div",{class:s.a.cardHeader},p,m),Object(o.h)("div",{class:s.a.cardBody},"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt."),h))},e}(o.Component)},sJaT:function(t,e,a){"use strict";var n=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var a=arguments[e];for(var n in a)Object.prototype.hasOwnProperty.call(a,n)&&(t[n]=a[n])}return t},r=a("SpGf");Object.defineProperty(e,"__esModule",{value:!0}),e.default=e.Card=e.CardMediaContent=e.CardActionIcon=e.CardActionButtons=e.CardActionIcons=e.CardActionButton=e.CardMedia=e.CardActions=void 0;var o=r(a("J5U+")),u=r(a("0fcM")),i=r(a("P8NW")),c=r(a("0421")),d=r(a("UJE0")),l=r(a("d4H2")),s=a("KM04"),f=r(a("uc5p")),p=r(a("7/cg")),m=r(a("MeGi")),h=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__actions",t.mdcProps=["full-bleed"],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return(0,s.h)("div",n({},t),this.props.children)}}]),e}(f.default);e.CardActions=h;var v=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__media",t.mdcProps=["square","16-9"],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return t.sixteenByNine&&(t.className="mdc-card__media--16-9"),(0,s.h)("div",n({},t),this.props.children)}}]),e}(f.default);e.CardMedia=v;var y=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__action",t.mdcProps=[],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return(0,s.h)("button",n({className:"mdc-button mdc-card__action--button"},t,{ref:this.setControlRef}),t.children)}}]),e}(p.default);e.CardActionButton=y;var _=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__action-icons",t.mdcProps=[],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return(0,s.h)("div",n({},t),this.props.children)}}]),e}(f.default);e.CardActionIcons=_;var b=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__action-buttons",t}return(0,l.default)(e,t),e}(_);e.CardActionButtons=b;var O=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__action",t.mdcProps=[],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return t.className?t.className+=" mdc-card__action--icon":t.className="mdc-card__action--icon",(0,o.default)((0,d.default)(e.prototype),"materialDom",this).call(this,t)}}]),e}(m.default);e.CardActionIcon=O;var C=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card__media-content",t.mdcProps=[],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return(0,s.h)("div",n({},t),this.props.children)}}]),e}(f.default);e.CardMediaContent=C;var A=function(t){function e(){var t;return(0,u.default)(this,e),t=(0,c.default)(this,(0,d.default)(e).apply(this,arguments)),t.componentName="card",t.mdcProps=["outlined"],t}return(0,l.default)(e,t),(0,i.default)(e,[{key:"materialDom",value:function(t){return(0,s.h)("div",n({},t),this.props.children)}}]),e}(f.default);e.Card=A;var j=function(t){function e(){return(0,u.default)(this,e),(0,c.default)(this,(0,d.default)(e).apply(this,arguments))}return(0,l.default)(e,t),e}(A);e.default=j,j.Actions=h,j.ActionButtons=b,j.ActionButton=y,j.ActionIcons=_,j.ActionIcon=O,j.Media=v,j.CardMediaContent=C}});
//# sourceMappingURL=route-home.chunk.970ea.js.map