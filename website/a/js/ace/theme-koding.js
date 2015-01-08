ace.define("ace/theme/koding",["require","exports","module","ace/lib/dom"],function(acequire,exports,module){exports.isDark=true;exports.cssClass="ace-koding";exports.cssText=".ace-koding {background-color: #1B2224;color: #93A1A1}.ace-koding .ace_gutter {background: #1B2224;color: #5e7074;}.ace-koding .ace_print-margin {width: 1px;background: #252E30;}.ace-koding .ace_entity.ace_other.ace_attribute-name,.ace-koding .ace_storage {color: #93A1A1}.ace-koding .ace_cursor {border-left: 2px solid #e9e7e1;}.ace-koding .ace_overwrite-cursors .ace_cursor {border-left: 0px;border-bottom: 1px solid #e9e7e1;}.ace-koding .ace_marker-layer .ace_active-line,.ace-koding .ace_marker-layer .ace_selection {background: rgba(255, 255, 255, 0.1)}.ace-koding.ace_multiselect .ace_selection.ace_start {box-shadow: 0 0 3px 0px #002B36;border-radius: 2px}.ace-koding .ace_marker-layer .ace_step {background: rgb(102, 82, 0)}.ace-koding .ace_marker-layer .ace_bracket {margin: -1px 0 0 -1px;border: 1px solid rgba(147, 161, 161, 0.50)}.ace-koding .ace_gutter-active-line {background-color: #0d3440}.ace-koding .ace_marker-layer .ace_selected-word {border: 1px solid #073642}.ace-koding .ace_invisible {color: rgba(147, 161, 161, 0.50)}.ace-koding .ace_keyword,.ace-koding .ace_meta,.ace-koding .ace_support.ace_class,.ace-koding .ace_support.ace_type {color: #859900}.ace-koding .ace_constant.ace_character,.ace-koding .ace_constant.ace_other {color: #CB4B16}.ace-koding .ace_constant.ace_language {color: #B58900}.ace-koding .ace_constant.ace_numeric {color: #D33682}.ace-koding .ace_fold {background-color: #268BD2;border-color: #93A1A1}.ace-koding .ace_entity.ace_name.ace_function,.ace-koding .ace_entity.ace_name.ace_tag,.ace-koding .ace_support.ace_function,.ace-koding .ace_variable,.ace-koding .ace_variable.ace_language {color: #268BD2}.ace-koding .ace_string {color: #2AA198}.ace-koding .ace_string.ace_regexp {color: #D30102}.ace-koding .ace_comment {font-style: italic;color: #657B83}.ace-koding .ace_indent-guide {background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAACCAYAAACZgbYnAAAAEklEQVQImWNg0Db1ZVCxc/sPAAd4AlUHlLenAAAAAElFTkSuQmCC) right repeat-y;}";var dom=acequire("../lib/dom");dom.importCssString(exports.cssText,exports.cssClass)});