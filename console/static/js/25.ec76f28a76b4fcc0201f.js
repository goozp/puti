webpackJsonp([25],{"+qOE":function(t,e,s){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=s("+aXT"),i={data:function(){return{settingType:"property",form:{site_description:"",site_keywords:"",footer_copyright:""}}},created:function(){this.setTitle(),this.getOptions()},methods:{setTitle:function(){document.title=this.$t("route."+this.$route.meta.title)+" | Puti"},getQuery:function(){return{settingType:this.settingType}},getOptions:function(){var t=this;o.a(this.getQuery()).then(function(e){var s=e.data;t.form.site_description=s.site_description,t.form.site_keywords=s.site_keywords,t.form.footer_copyright=s.footer_copyright})},saveSetting:function(){var t=this;o.b(this.getQuery(),this.form).then(function(e){t.$message({message:t.$t("common.operationSucceeded"),type:"success",duration:2e3}),t.getOptions()})},resetSetting:function(){this.getOptions()}}},r={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("el-card",{staticClass:"box-card setting-form",attrs:{shadow:"hover"}},[s("div",{staticClass:"setting-form-header clearfix",attrs:{slot:"header"},slot:"header"},[s("span",{staticClass:"setting-header-title"},[t._v(t._s(t.$t("option.siteSetting")))]),t._v(" "),s("span",{staticClass:"setting-header-subtitle"},[t._v(t._s(t.$t("option.siteSettingProperty")))])]),t._v(" "),s("div",{},[s("el-form",{ref:"form",attrs:{model:t.form,"label-width":"200px"}},[s("div",{staticClass:"setting-form-body"},[s("el-row",[s("el-col",{attrs:{span:14,offset:2}},[s("el-form-item",{attrs:{label:t.$t("option.siteDescription"),prop:"site_description"}},[s("el-input",{attrs:{type:"textarea",rows:"5"},model:{value:t.form.site_description,callback:function(e){t.$set(t.form,"site_description",e)},expression:"form.site_description"}}),t._v(" "),s("p",{staticClass:"setting-form-desc"},[t._v(t._s(t.$t("option.siteDescriptionDesc")))])],1),t._v(" "),s("el-form-item",{attrs:{label:t.$t("option.siteKeywords"),prop:"site_keywords"}},[s("el-input",{attrs:{type:"textarea",rows:"5"},model:{value:t.form.site_keywords,callback:function(e){t.$set(t.form,"site_keywords",e)},expression:"form.site_keywords"}}),t._v(" "),s("p",{staticClass:"setting-form-desc"},[t._v(t._s(t.$t("option.siteKeywordsDesc")))])],1),t._v(" "),s("el-form-item",{attrs:{label:t.$t("option.siteFooter"),prop:"footer_copyright"}},[s("el-input",{attrs:{type:"textarea",rows:"10"},model:{value:t.form.footer_copyright,callback:function(e){t.$set(t.form,"footer_copyright",e)},expression:"form.footer_copyright"}}),t._v(" "),s("p",{staticClass:"setting-form-desc"},[t._v(t._s(t.$t("option.siteFooterDesc")))])],1)],1)],1),t._v(" "),s("el-row",{staticClass:"setting-form-footer"},[s("el-col",{attrs:{span:14,offset:6}},[s("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.saveSetting}},[t._v(t._s(t.$t("option.siteSettingSave")))]),t._v(" "),s("el-button",{attrs:{size:"small"},on:{click:t.resetSetting}},[t._v(t._s(t.$t("common.reset")))])],1)],1)],1)])],1)])},staticRenderFns:[]},n=s("VU/8")(i,r,!1,null,null,null);e.default=n.exports}});