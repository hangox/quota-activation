package management

import (
	"net/http"
)

// allow: SIZE_OK - 内嵌静态资源页以单 HTML 载荷交给 CPA resources 路由；拆分不会降低运行时职责。
func (h *Handler) handlePage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<!doctype html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>配额唤醒 · quota-activation</title>
  <script>
    // 首次绘制前标出内嵌模式，避免顶部安全带跳动（与共享外壳的判定一致）。
    document.documentElement.dataset.embedded = String(window.self !== window.top);
  </script>
  <style id="plugin-page-shell">
/* 由 assets/plugin-page-shell.css 生成 · 请勿手改；改真源后重跑 scripts/build-plugin-pages.mjs */
/*
 * 插件页面共享外壳样式（唯一真源）。
 * 由 scripts/build-plugin-pages.mjs 在构建期注入到各插件页的 <style id="plugin-page-shell"> 块。
 * 不要直接修改生成后的 Go 文件内联副本；改这里后重跑生成脚本。
 *
 * 设计意图：让插件页面看起来像 CPAMC 原生工作区 —— 正文全宽、顶部纵向让开宿主浮动工具栏、
 * 使用宿主暖灰/纯白 token，并把「页头 → 摘要 chips → Tab → 工作区」的层次固定下来。
 * 只承载跨插件复用的 token 与基础组件；页面专属样式放在各插件页自己的 <style> 里。
 */

:root {
  color-scheme: light;
  /* 主题 token（white 为默认） */
  --bg-secondary: #ffffff;
  --bg-primary: #ffffff;
  --bg-tertiary: #f6f5f2;
  --text-primary: #2d2a26;
  --text-secondary: #6d6760;
  --text-muted: #9a948c;
  --border-color: #e5e5e5;
  --border-hover: #cccccc;
  --primary-color: #8b8680;
  --primary-hover: #7f7a74;
  --primary-contrast: #ffffff;
  --primary-soft: color-mix(in srgb, var(--primary-color) 12%, transparent);
  --shadow: 0 1px 2px rgb(0 0 0 / 8%);
  /* 语义状态色：只表达状态，配文字/图标，不当作系列色 */
  --success: #1a7f47;
  --success-bg: #eef8f1;
  --success-line: #cfe9d8;
  --danger: #b42318;
  --danger-bg: #fdf2f1;
  --danger-line: #f3d2ce;
  --warn: #95610a;
  --warn-bg: #fdf6e5;
  --warn-line: #efdfb4;
  --info: #40607f;
  --info-bg: #eef3f8;
  --info-line: #d3e0ec;
  --radius: 12px;
  --radius-sm: 8px;
  --radius-pill: 999px;
  /* 控件高度 token：手机断点只抬升这些变量即可满足 44px，避免「元素选择器 < 类选择器」
     的优先级退化（例如 select.input 的硬编码高度压过媒体查询里的 select{min-height:44px}）。 */
  --control-h: 34px;
  --control-h-input: 38px;
  --control-h-select: 34px;
  --mono: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
  /* 顶部安全带：裸页小，内嵌时由 JS/内联标记抬到宿主浮层之下 */
  --page-top: 24px;
  --page-gutter: clamp(24px, 3vw, 48px);
}

:root[data-embedded="true"] { --page-top: 94px; }

:root[data-theme="light"] {
  --bg-secondary: #faf9f5;
  --bg-primary: #f0eee8;
  --bg-tertiary: #e9e6df;
  --border-color: #e3e1db;
  --border-hover: #cecac4;
}

:root[data-theme="dark"] {
  color-scheme: dark;
  --bg-secondary: #151412;
  --bg-primary: #1d1b18;
  --bg-tertiary: #262320;
  --text-primary: #f6f4f1;
  --text-secondary: #c9c3bb;
  --text-muted: #928b82;
  --border-color: #3a3530;
  --border-hover: #5a544d;
  --primary-hover: #9a948e;
  --success: #5bc98a;
  --success-bg: #14261b;
  --success-line: #23452f;
  --danger: #f08a80;
  --danger-bg: #2a1614;
  --danger-line: #4a2621;
  --warn: #e0b34a;
  --warn-bg: #2a2210;
  --warn-line: #47381a;
  --info: #9fb4cc;
  --info-bg: #16202b;
  --info-line: #2a3a4a;
  --shadow: 0 1px 3px rgb(0 0 0 / 30%);
}

* { box-sizing: border-box; }

[hidden] { display: none !important; }

body {
  margin: 0;
  background: var(--bg-secondary);
  color: var(--text-primary);
  font: 14px/1.5 -apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
  -webkit-font-smoothing: antialiased;
}

button, input, select { font: inherit; }
button { cursor: pointer; }
button:disabled { opacity: .6; cursor: not-allowed; }

button:focus-visible,
input:focus-visible,
select:focus-visible,
.tab:focus-visible,
.table-scroll:focus-visible,
[tabindex]:focus-visible {
  outline: 2px solid var(--primary-color);
  outline-offset: 2px;
}

/* 页面骨架：全宽 + 纵向安全带。 */
.page {
  width: 100%;
  min-width: 0;
  padding: var(--page-top) var(--page-gutter) 40px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.page-header { display: flex; flex-direction: column; gap: 12px; }
.header-row { display: flex; justify-content: space-between; align-items: flex-end; gap: 16px; }
.title-area { min-width: 0; }
h1 { margin: 0; font-size: 30px; line-height: 1.2; font-weight: 700; overflow-wrap: anywhere; }
h2 { margin: 0; font-size: 22px; line-height: 1.3; font-weight: 600; }
h3 { margin: 0; font-size: 16px; line-height: 1.35; font-weight: 600; }
.description { margin: 6px 0 0; color: var(--text-secondary); }
.actions, .chips { display: flex; flex-wrap: wrap; align-items: center; gap: 8px; }
.actions { flex-shrink: 0; }
.chip {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  min-height: var(--control-h);
  padding: 5px 12px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  background: var(--bg-tertiary);
  color: var(--text-secondary);
  font-size: 13px;
}
.chip strong { color: var(--text-primary); font-weight: 600; }

/* 按钮 */
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  min-height: var(--control-h);
  padding: 6px 14px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border-color);
  background: var(--bg-primary);
  color: var(--text-primary);
  transition: background-color 150ms ease, border-color 150ms ease;
}
.btn:hover:not(:disabled) { background: var(--bg-tertiary); border-color: var(--border-hover); }
.btn-primary { background: var(--primary-color); border-color: transparent; color: var(--primary-contrast); }
.btn-primary:hover:not(:disabled) { background: var(--primary-hover); border-color: transparent; }
.btn-danger { background: var(--danger); border-color: transparent; color: #fff; }
.btn-danger:hover:not(:disabled) { background: color-mix(in srgb, var(--danger) 88%, #000); border-color: transparent; }

/* 视图 Tab（页级） */
.tabs { display: flex; flex-wrap: wrap; gap: 4px; }
.tab {
  display: inline-flex;
  align-items: center;
  min-height: var(--control-h);
  padding: 9px 14px;
  border: 1px solid transparent;
  border-radius: var(--radius-sm);
  background: transparent;
  color: var(--text-primary);
}
.tab:hover { background: var(--bg-tertiary); border-color: var(--border-color); }
.tab[aria-selected="true"] {
  color: var(--text-primary);
  border-color: color-mix(in srgb, var(--primary-color) 35%, transparent);
  background: var(--primary-soft);
  font-weight: 600;
}

/* 面板 */
.panel {
  min-width: 0;
  padding: 20px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius);
  background: var(--bg-primary);
  box-shadow: var(--shadow);
}
.panel-header { display: flex; flex-direction: column; gap: 8px; margin-bottom: 16px; }
.panel-note { margin: 0; color: var(--text-secondary); font-size: 13px; }

/* 空 / 加载 / 错误状态：不伪装成正常数据 */
.empty {
  padding: 32px;
  border: 1px dashed var(--border-color);
  border-radius: var(--radius-sm);
  text-align: center;
  color: var(--text-secondary);
}
.empty p { margin: 0 0 12px; }
.state-inline { color: var(--text-secondary); }
.state-error {
  margin: 0;
  padding: 10px 14px;
  border: 1px solid var(--danger-line);
  border-radius: var(--radius-sm);
  background: var(--danger-bg);
  color: var(--danger);
  font-size: 13px;
}

/* 表格：横向滚动只发生在容器内 */
.table-scroll { overflow-x: auto; border: 1px solid var(--border-color); border-radius: var(--radius-sm); }
table { width: 100%; min-width: 640px; border-collapse: collapse; font-size: 13px; }
th, td { padding: 11px 12px; border-bottom: 1px solid var(--border-color); text-align: left; vertical-align: top; }
th { background: var(--bg-tertiary); color: var(--text-secondary); font-weight: 500; white-space: nowrap; }
tr:last-child td { border-bottom: 0; }
td.num, th.num { text-align: right; font-variant-numeric: tabular-nums; }
td.mono { font-family: var(--mono); font-size: 12px; color: var(--text-secondary); }
.empty-cell { text-align: center; color: var(--text-secondary); padding: 32px; }

/* 状态徽标：颜色服务语义，同时带文字 */
.badge {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 2px 9px;
  border: 1px solid transparent;
  border-radius: var(--radius-pill);
  font-size: 12px;
  font-weight: 600;
  white-space: nowrap;
}
.badge.ok { background: var(--success-bg); color: var(--success); border-color: var(--success-line); }
.badge.warn { background: var(--warn-bg); color: var(--warn); border-color: var(--warn-line); }
.badge.bad { background: var(--danger-bg); color: var(--danger); border-color: var(--danger-line); }
.badge.info { background: var(--info-bg); color: var(--info); border-color: var(--info-line); }

/* 数值卡：明确单位/口径；无数据显示「—」而不是 0 */
.tiles { display: grid; gap: 12px; }
.tile {
  display: flex;
  flex-direction: column;
  gap: 5px;
  min-width: 0;
  padding: 14px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius);
  background: var(--bg-primary);
  box-shadow: var(--shadow);
}
.tile .tile-label { color: var(--text-secondary); font-size: 12px; }
.tile .tile-value { font-size: 22px; font-weight: 700; letter-spacing: -.2px; font-variant-numeric: tabular-nums; }
.tile .tile-note { color: var(--text-muted); font-size: 12px; }

/* 表单控件 */
.field-label { display: block; margin: 0 0 6px; color: var(--text-secondary); font-size: 13px; font-weight: 600; }
/* 表单控件：高度全部来自 token，手机断点只改 token 即可兜底 44px */
.input, input { min-height: var(--control-h-input); }
select, select.input { min-height: var(--control-h-select); }
.input, select.input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  background: var(--bg-primary);
  color: var(--text-primary);
}
select.input { width: auto; padding: 6px 10px; }
.toolbar { display: flex; flex-wrap: wrap; align-items: center; gap: 8px; }

@media (max-width: 768px) {
  :root {
    --page-top: 16px; --page-gutter: 16px;
    /* 主要控件（按钮/输入/下拉）手机兜底 ≥44px：只抬升 token，不写元素级 min-height，
       因此不会被 .input / select.input 这类更高优先级的类选择器覆盖。 */
    --control-h: 44px; --control-h-input: 44px; --control-h-select: 44px;
  }
  :root[data-embedded="true"] { --page-top: 112px; }
  .page { gap: 16px; padding-bottom: 24px; }
  .header-row { flex-direction: column; align-items: flex-start; }
  .panel { padding: 16px; }
  .empty { padding: 24px 16px; }
  .chip { padding: 10px 12px; }
  h1 { font-size: 24px; }
  h2 { font-size: 19px; }
}

@media (prefers-reduced-motion: reduce) { * { transition: none !important; } }
</style>
  <style>
    /* 页面专属样式：共享 token/基础组件来自外壳，这里只放 quota-activation 特有结构。 */
    .page-status { min-height: 22px; margin: 0; color: var(--text-secondary); font-size: 13px; white-space: pre-wrap; }
    .page-status:empty { display: none; }

    .gate-actions { margin-top: 14px; }

    .provider-tabs { gap: 6px; }
    .provider-tabs .provider-tab {
      min-height: var(--control-h); padding: 6px 14px; border: 1px solid var(--border-color);
      border-radius: var(--radius-pill); background: var(--bg-primary); color: var(--text-secondary);
    }
    .provider-tabs .provider-tab:hover { border-color: var(--border-hover); color: var(--text-primary); }
    .provider-tabs .provider-tab[aria-checked="true"] {
      background: var(--primary-color); border-color: transparent; color: var(--primary-contrast); font-weight: 600;
    }

    .select-shell { margin: 0 0 4px; border: 1px solid var(--border-color); border-radius: var(--radius-sm); background: var(--bg-primary); overflow: hidden; }
    .credential-list { max-height: 240px; overflow: auto; padding: 8px; display: flex; flex-direction: column; gap: 4px; }
    .credential-item { display: flex; align-items: center; gap: 10px; margin: 0; padding: 8px 10px; border-radius: var(--radius-sm); color: var(--text-primary); cursor: pointer; }
    .credential-item:hover { background: var(--bg-tertiary); }
    .credential-item.disabled { color: var(--text-muted); cursor: not-allowed; }
    .credential-item.disabled:hover { background: transparent; }
    .credential-item input { width: auto; margin: 0; }
    .credential-item.disabled input { pointer-events: none; }

    .model-menu { position: relative; }
    .model-menu-button {
      width: 100%; min-height: var(--control-h-input); display: flex; align-items: center; justify-content: space-between; gap: 12px;
      border: 1px solid var(--border-color); border-radius: var(--radius-sm); padding: 0 14px;
      background: var(--bg-primary); color: var(--text-primary); text-align: left;
    }
    .model-menu-panel {
      position: absolute; left: 0; right: 0; top: calc(100% + 6px); z-index: 40; display: none;
      max-height: 220px; overflow: auto; border: 1px solid var(--border-color); border-radius: var(--radius-sm);
      background: var(--bg-primary); box-shadow: 0 16px 36px rgb(0 0 0 / 14%); padding: 6px;
    }
    .model-menu-panel.open { display: block; }
    .model-option { width: 100%; min-height: var(--control-h); display: block; border: 0; border-radius: var(--radius-sm); padding: 8px 12px; background: transparent; color: var(--text-primary); text-align: left; }
    .model-option:hover, .model-option.active { background: var(--bg-tertiary); }
    .model-option.disabled { color: var(--text-muted); cursor: not-allowed; opacity: .55; }
    .model-option.disabled:hover { background: transparent; }

    .confirm-box {
      margin-top: 16px; padding: 16px; border: 1px solid var(--border-color);
      border-radius: var(--radius-sm); background: var(--bg-tertiary);
    }
    .confirm-box h3 { margin-bottom: 8px; }
    .confirm-list { margin: 0 0 12px; padding-left: 18px; color: var(--text-secondary); font-size: 13px; }
    .confirm-list li { margin-bottom: 4px; }
    .confirm-cost { margin: 0 0 12px; color: var(--warn); font-size: 13px; }

    .records-header { display: flex; flex-wrap: wrap; align-items: flex-end; justify-content: space-between; gap: 12px; margin-bottom: 16px; }
    .records-hint { margin: 4px 0 0; color: var(--text-secondary); font-size: 13px; }
    .records-list { display: grid; gap: 12px; }
    .record-card { border: 1px solid var(--border-color); border-radius: var(--radius); padding: 14px 16px; display: grid; gap: 10px; background: var(--bg-primary); }
    .record-card.warn { border-color: var(--warn-line); background: var(--warn-bg); }
    .record-card.fail { border-color: var(--danger-line); background: var(--danger-bg); }
    .record-head { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 10px; }
    .record-meta { display: flex; flex-wrap: wrap; gap: 8px; align-items: center; }
    .record-time { font-size: 12px; color: var(--text-secondary); font-weight: 600; }
    .record-badge { display: inline-flex; align-items: center; padding: 2px 10px; border-radius: var(--radius-pill); font-size: 12px; font-weight: 650; }
    .record-badge.manual { background: var(--info-bg); color: var(--info); }
    .record-badge.auto { background: var(--success-bg); color: var(--success); }
    .record-badge.scan { background: var(--primary-soft); color: var(--text-primary); }
    .record-stats { display: flex; flex-wrap: wrap; gap: 8px; }
    .record-stat { display: inline-flex; align-items: center; gap: 4px; border-radius: var(--radius-pill); padding: 4px 10px; background: var(--bg-tertiary); border: 1px solid var(--border-color); font-size: 12px; font-weight: 650; color: var(--text-secondary); }
    .record-providers { display: grid; grid-template-columns: repeat(3, minmax(0, 1fr)); gap: 8px; }
    .record-provider { display: grid; gap: 4px; padding: 10px 12px; border-radius: var(--radius-sm); background: var(--bg-tertiary); border: 1px solid var(--border-color); min-width: 0; }
    .record-provider strong { font-size: 13px; color: var(--text-primary); }
    .record-provider span { font-size: 12px; color: var(--text-secondary); overflow-wrap: anywhere; }
    .record-provider .prov-err { color: var(--danger); }
    .record-provider .prov-warn { color: var(--warn); }
    .record-message { font-size: 13px; overflow-wrap: anywhere; line-height: 1.5; }
    .record-message.error { color: var(--danger); }
    .record-message.warn { color: var(--warn); }
    .empty-state { padding: 28px 12px; border: 1px dashed var(--border-color); border-radius: var(--radius-sm); text-align: center; color: var(--text-secondary); font-size: 14px; }

    .help-content { line-height: 1.6; color: var(--text-primary); }
    .help-content h3 { margin: 20px 0 8px; }
    .help-content h3:first-child { margin-top: 0; }
    .help-content ul { padding-left: 20px; margin: 0 0 16px; }
    .help-content li { margin-bottom: 6px; color: var(--text-secondary); }
    .help-content code { background: var(--bg-tertiary); border-radius: 6px; padding: 2px 6px; font-family: var(--mono); font-size: 13px; }
    .help-content pre { background: var(--bg-tertiary); padding: 12px; border-radius: var(--radius-sm); overflow-x: auto; font-family: var(--mono); font-size: 13px; margin: 8px 0; }

    @media (max-width: 720px) { .record-providers { grid-template-columns: 1fr; } }
    @media (max-width: 768px) {
      /* 控件高度已由共享 token 抬到 44px；这里只补非 token 化的行级点击区。 */
      .credential-item { min-height: 44px; }
    }
  </style>
</head>
<body>
<main class="page">
  <header class="page-header">
    <div class="header-row">
      <div class="title-area">
        <h1 data-i18n="title">配额唤醒</h1>
        <p class="description" data-i18n="pageDesc">quota-activation 插件：查看凭证并手动触发配额唤醒，执行记录与配置说明在此查看。</p>
      </div>
      <div class="actions">
        <button id="languageMenuButton" type="button" class="btn" onclick="switchLanguage()" aria-label="Switch language" hidden>
          <span data-i18n="languageLabel">语言</span>
        </button>
        <span class="chip"><span data-i18n="pluginLabel">插件</span> · quota-activation</span>
        <span class="chip"><strong>v0.0.7</strong></span>
      </div>
    </div>
    <div class="chips" id="headerChips" hidden>
      <span class="chip"><span data-i18n="chipCredentials">凭证</span> <strong id="chipCredCount">—</strong></span>
      <span class="chip"><span data-i18n="chipAuto">自动唤醒</span> <strong id="chipAuto">—</strong></span>
    </div>
  </header>

  <section class="panel" id="loginGate">
    <div class="panel-header">
      <h2 data-i18n="gateTitle">管理密钥</h2>
      <p class="panel-note" data-i18n="gateDesc">插件资源页匿名可打开；查看凭证与执行唤醒需要 CPA 管理密钥，密钥只在本次请求中使用。</p>
    </div>
    <label class="field-label" for="managementKey" data-i18n="managementKeyLabel">CPA 管理密钥</label>
    <input class="input" id="managementKey" type="password" autocomplete="current-password" data-i18n-placeholder="managementKeyPlaceholder" placeholder="输入 CPA 管理密钥">
    <div class="actions gate-actions">
      <button class="btn btn-primary" type="button" data-i18n="verifyKey" onclick="verifyManagementKey(this)">验证密钥</button>
    </div>
  </section>

  <div id="appShell" hidden>
    <div class="tabs" role="tablist" aria-orientation="horizontal" aria-label="配额唤醒视图">
      <button class="tab" type="button" role="tab" id="manualTabButton" aria-controls="manualPanel" aria-selected="true" tabindex="0" data-tab="manual" onclick="switchTab('manual')" data-i18n="manualTab">手动唤醒</button>
      <button class="tab" type="button" role="tab" id="recordsTabButton" aria-controls="recordsPanel" aria-selected="false" tabindex="-1" data-tab="records" onclick="switchTab('records')" data-i18n="recordsTab">执行记录</button>
      <button class="tab" type="button" role="tab" id="helpTabButton" aria-controls="helpPanel" aria-selected="false" tabindex="-1" data-tab="help" onclick="switchTab('help')" data-i18n="helpTab">帮助</button>
    </div>
    <p id="managementToast" class="page-status" role="status" aria-live="polite"></p>

    <section class="panel" role="tabpanel" id="manualPanel" aria-labelledby="manualTabButton" data-panel="manual">
      <div class="panel-header">
        <h2 data-i18n="manualTab">手动唤醒</h2>
        <p class="panel-note" data-i18n="manualNote">选择凭证与唤醒模型后手动触发；执行前会再次确认，执行中不可重复提交。</p>
      </div>
      <label class="field-label" for="credentialSearch" data-i18n="credential">凭证</label>
      <input class="input" id="credentialSearch" type="search" data-i18n-placeholder="searchPlaceholder" placeholder="搜索凭证..." oninput="renderAuthFiles()" style="margin-bottom:10px">
      <div id="providerTabs" class="tabs provider-tabs" role="radiogroup" aria-label="按提供商筛选凭证"></div>
      <div id="credentialSelect" class="select-shell credential-list" role="group" aria-label="credentials"></div>
      <label class="field-label" for="modelMenuButton" data-i18n="model">唤醒模型</label>
      <div class="model-menu" id="modelSelect">
        <button id="modelMenuButton" type="button" class="model-menu-button" onclick="toggleModelMenu()" aria-haspopup="listbox" aria-expanded="false">
          <span id="modelMenuLabel">—</span>
          <span aria-hidden="true">▾</span>
        </button>
        <div id="modelMenuPanel" class="model-menu-panel" role="listbox"></div>
      </div>
      <div class="actions" style="margin-top:16px">
        <button id="refreshCredentialsButton" type="button" class="btn" data-i18n="refresh" onclick="loadAuthFiles(this)">刷新凭证</button>
        <button id="manualActivateButton" type="button" class="btn btn-primary" data-i18n="activate" onclick="requestActivation(this)">手动触发唤醒</button>
      </div>
      <div id="confirmBox" class="confirm-box" hidden>
        <h3 data-i18n="confirmTitle">确认执行唤醒</h3>
        <ul class="confirm-list" id="confirmTargets"></ul>
        <p class="confirm-cost" data-i18n="confirmCost">执行会产生一次模型请求，可能消耗所选凭证的模型额度。</p>
        <div class="actions">
          <button id="confirmCancelButton" type="button" class="btn" data-i18n="confirmCancel" onclick="cancelActivation()">取消</button>
          <button id="confirmRunButton" type="button" class="btn btn-primary" data-i18n="confirmRun" onclick="confirmActivation(this)">确认执行</button>
        </div>
      </div>
    </section>

    <section class="panel" role="tabpanel" id="recordsPanel" aria-labelledby="recordsTabButton" data-panel="records" hidden>
      <div class="panel-header">
        <h2 data-i18n="recordsTitle">执行记录</h2>
        <p class="panel-note" data-i18n="recordsHint">进程内最近 5 次（重启清空）</p>
      </div>
      <div class="records-header">
        <div class="records-header-text">
          <p class="records-hint" data-i18n="recordsScopeNote">仅显示进程内摘要，取自已脱敏的诊断数据。</p>
        </div>
        <button id="refreshRecordsButton" type="button" class="btn" data-i18n="refreshRecords" onclick="loadRunHistory(this)">刷新记录</button>
      </div>
      <div id="recordsList" class="records-list"></div>
    </section>

    <section class="panel" role="tabpanel" id="helpPanel" aria-labelledby="helpTabButton" data-panel="help" hidden>
      <div class="panel-header">
        <h2 data-i18n="helpTab">帮助</h2>
      </div>
      <div class="help-content">
        <h3 data-i18n="helpConfigTitle">配置字段说明</h3>
        <p data-i18n="helpConfigDesc">配置 quota-activation 插件所需的字段如下：</p>
        <ul>
          <li><code>auto_activate</code>: <span data-i18n="helpAutoActivate">是否开启自动配额唤醒。</span></li>
          <li><code>enable_before_activation</code>: <span data-i18n="helpEnableBefore">是否在唤醒前自动启用已被禁用的凭证。</span></li>
          <li><code data-optional="true">scan_interval</code>: <span data-i18n="helpScanInterval">可选；自动扫描间隔，单位分钟，填写纯数字即可（默认 30）。</span></li>
          <li><code data-optional="true">activation_request_timeout</code>: <span data-i18n="helpActivationTimeout">可选；唤醒请求超时，单位秒，填写纯数字即可（默认 60）。</span></li>
          <li><code data-optional="true">max_concurrency</code>: <span data-i18n="helpMaxConcurrency">可选；自动扫描与激活的并发上限（worker 池大小），默认 1。</span></li>
          <li><code data-optional="true">activation_prompt</code>: <span data-i18n="helpActivationPrompt">可选；唤醒提示词。</span></li>
        </ul>
        <h3 data-i18n="helpActivationModelsTitle">自动唤醒模型字段</h3>
        <p data-i18n="helpActivationModelsDesc">直接填写模型名称，不需要填写 JSON 对象：</p>
        <ul>
          <li><code>activation_models.codex.models</code>: <span data-i18n="helpCodexModels">Codex 自动唤醒模型名称，例如 <code>gpt-5-mini</code>。</span></li>
          <li><code>activation_models.antigravity.models_group</code>: <span data-i18n="helpAntigravityGroup">Antigravity 自动唤醒模型组，可选 <code>gemini</code> 或 <code>claude_gpt</code>。</span></li>
          <li><code>activation_models.antigravity.models</code>: <span data-i18n="helpAntigravityModels">当前 Antigravity 模型组的模型名称。</span></li>
        </ul>
      </div>
    </section>
  </div>
</main>
<script id="plugin-page-shell">
/* 由 assets/plugin-page-shell.js 生成 · 请勿手改；改真源后重跑 scripts/build-plugin-pages.mjs */
/*
 * 插件页面共享适配逻辑（唯一真源）。
 * 由 scripts/build-plugin-pages.mjs 在构建期注入到各插件页的 <script id="plugin-page-shell"> 块。
 *
 * 只做两件事：把 iframe 缺失的宿主顶部安全带补回来，并只读同步宿主可见视觉 token。
 * 隐私边界：仅读取同源父文档的 DOM 属性与 computed style；不读 Cookie、localStorage、
 * sessionStorage、认证 store，也不接收任意来源的 postMessage。
 * 跨源时不能访问父文档，回退到保守的 94px/112px 安全带与系统主题。
 */
(function () {
  var root = document.documentElement;
  var embedded = window.self !== window.top;
  root.dataset.embedded = String(embedded);

  var systemTheme = window.matchMedia('(prefers-color-scheme: dark)');
  var VISUAL_TOKENS = [
    '--bg-secondary', '--bg-primary', '--bg-tertiary',
    '--text-primary', '--text-secondary', '--text-muted',
    '--border-color', '--border-hover',
    '--primary-color', '--primary-hover', '--primary-contrast', '--shadow'
  ];
  var hostDocument = null;
  try { if (embedded) { hostDocument = window.parent.document; } }
  catch (error) { hostDocument = null; }

  function sameOriginHost() {
    try { void window.parent.document.documentElement; return hostDocument; }
    catch (error) { return null; }
  }

  function syncAppearance() {
    var host = sameOriginHost();
    if (!host) {
      root.dataset.theme = systemTheme.matches ? 'dark' : 'white';
      return;
    }
    var hostRoot = host.documentElement;
    var theme = hostRoot.dataset.theme || 'light';
    root.dataset.theme = (theme === 'white' || theme === 'dark') ? theme : 'light';

    var hostStyle = window.parent.getComputedStyle(hostRoot);
    for (var i = 0; i < VISUAL_TOKENS.length; i++) {
      var value = hostStyle.getPropertyValue(VISUAL_TOKENS[i]).trim();
      if (value) { root.style.setProperty(VISUAL_TOKENS[i], value); }
    }

    var frame = window.frameElement;
    var frameTop = frame ? frame.getBoundingClientRect().top : 0;
    var actions = host.querySelector('.header-actions');
    var mobile = window.parent.innerWidth <= 768;
    var headerHeight = parseFloat(hostStyle.getPropertyValue('--header-height')) || 80;
    var fallbackTop = mobile ? headerHeight + 32 : 94;
    var bottomInFrame = actions ? actions.getBoundingClientRect().bottom - frameTop : 0;
    root.style.setProperty('--page-top', Math.ceil(Math.max(0, bottomInFrame + 18, fallbackTop - frameTop)) + 'px');
    var gutter = mobile ? 32 : Math.min(48, Math.max(20, window.parent.innerWidth * 0.03)) + 24;
    root.style.setProperty('--page-gutter', gutter + 'px');

    var hostLang = hostRoot.getAttribute('lang') || '';
    root.dataset.hostLang = hostLang;
  }

  syncAppearance();
  systemTheme.addEventListener('change', syncAppearance);

  var host = sameOriginHost();
  if (host) {
    window.parent.addEventListener('resize', syncAppearance);
    var themeObserver = new MutationObserver(syncAppearance);
    themeObserver.observe(host.documentElement, { attributes: true, attributeFilter: ['data-theme', 'style', 'lang'] });
    var layoutObserver = new ResizeObserver(syncAppearance);
    var actions = host.querySelector('.header-actions');
    if (actions) { layoutObserver.observe(actions); }
    if (window.frameElement) { layoutObserver.observe(window.frameElement); }
    window.addEventListener('pagehide', function () {
      themeObserver.disconnect();
      layoutObserver.disconnect();
      window.parent.removeEventListener('resize', syncAppearance);
      systemTheme.removeEventListener('change', syncAppearance);
    }, { once: true });
  }
})();
</script>
<script>
  const STATUS_PATH = "/v0/management/quota-activation/status";
  const AUTH_FILES_PATH = "/v0/management/quota-activation/auth-files";
  const AUTH_FILE_MODELS_PATH = "/v0/management/auth-files/models";
  const ACTIVATE_PATH = "/v0/management/quota-activation/activate";
  const DIAGNOSTICS_PATH = "/v0/management/quota-activation/diagnostics";
  const translations = {
    "zh-CN": {
      activate: "手动触发唤醒",
      activationDone: "手动唤醒请求已完成。",
      activationFailed: "手动唤醒失败，请稍后重试或检查凭证状态。",
      statusFailed: "唤醒失败",
      activationRunning: "唤醒中...",
      credential: "凭证",
      disabled: "已禁用",
      credentialDisabled: "请先启用凭证后再尝试唤醒",
      loadingCredentials: "刷新中...",
      keyFailed: "管理密钥验证失败，请检查后重试。",
      keyPassed: "管理密钥验证通过。",
      managementKeyPlaceholder: "输入 CPA 管理密钥",
      missingCredential: "请先选择凭证。",
      missingKey: "请先输入 CPA 管理密钥。",
      missingModel: "请先选择已启用的唤醒模型。",
      model: "唤醒模型",
      noAuthFiles: "未发现可唤醒凭证。",
      noModels: "该凭证没有启用的唤醒模型组。",
      refresh: "刷新凭证",
      statusBusy: "执行中",
      statusSkipped: "已跳过",
      statusSuccess: "成功",
      statusUnknown: "未知",
      title: "配额唤醒",
      usable: "可用",
      verifyKey: "验证密钥",
      verifyingKey: "验证中...",
      manualTab: "手动唤醒",
      helpTab: "帮助",
      searchPlaceholder: "搜索凭证...",
      languageLabel: "语言",
      noMatchingCredential: "没有匹配的凭证。",
      helpConfigTitle: "配置字段说明",
      helpConfigDesc: "配置 quota-activation 插件所需的字段如下：",
      helpActivationModelsTitle: "自动唤醒模型字段",
      helpActivationModelsDesc: "直接填写模型名称，不需要填写 JSON 对象：",
      helpAutoActivate: "是否开启自动配额唤醒。",
      helpEnableBefore: "是否在唤醒前自动启用已被禁用的凭证。",
      helpScanInterval: "可选；自动扫描间隔，单位分钟，填写纯数字即可（默认 30）。",
      helpActivationTimeout: "可选；唤醒请求超时，单位秒，填写纯数字即可（默认 60）。",
      helpMaxConcurrency: "可选；自动扫描与激活的并发上限（worker 池大小），默认 1。",
      helpActivationPrompt: "可选；唤醒提示词。",
      helpCodexModels: "Codex 自动唤醒模型名称，例如 gpt-5-mini。",
      helpAntigravityGroup: "Antigravity 自动唤醒模型组，可选 gemini 或 claude_gpt。",
      helpAntigravityModels: "当前 Antigravity 模型组的模型名称。",
      recordsTab: "执行记录",
      recordsTitle: "执行记录",
      recordsHint: "进程内最近 5 次（重启清空）",
      refreshRecords: "刷新记录",
      noRecords: "暂无执行记录",
      noRecordsAutoOff: "自动唤醒未启用。请在配置中将 auto_activate 设为 true；开启后每次扫描会写入摘要，也可手动唤醒产生记录。",
      triggerManual: "手动唤醒",
      triggerAuto: "自动唤醒",
      kindScan: "定时扫描",
      kindActivationAuto: "自动唤醒",
      kindActivationManual: "手动唤醒",
      providerBreakdown: "按凭证提供商",
      statAttempted: "尝试",
      statSucceeded: "成功",
      statFailed: "失败",
      statSkipped: "跳过",
      loadingRecords: "刷新中...",
      pageDesc: "quota-activation 插件：查看凭证并手动触发配额唤醒，执行记录与配置说明在此查看。",
      pluginLabel: "插件",
      gateTitle: "管理密钥",
      gateDesc: "插件资源页匿名可打开；查看凭证与执行唤醒需要 CPA 管理密钥，密钥只在本次请求中使用。",
      managementKeyLabel: "CPA 管理密钥",
      manualNote: "选择凭证与唤醒模型后手动触发；执行前会再次确认，执行中不可重复提交。",
      chipCredentials: "凭证",
      chipAuto: "自动唤醒",
      autoOn: "已开启",
      autoOff: "未开启",
      autoUnknown: "—",
      confirmTitle: "确认执行唤醒",
      confirmCost: "执行会产生一次模型请求，可能消耗所选凭证的模型额度。",
      confirmRun: "确认执行",
      confirmCancel: "取消",
      confirmTargetCredentials: "凭证",
      confirmTargetModel: "模型",
      recordsScopeNote: "仅显示进程内摘要，取自已脱敏的诊断数据。"
    },
    "en-US": {
      activate: "Trigger activation",
      activationDone: "Manual activation request completed.",
      activationFailed: "Manual activation failed. Try again later or check the credential status.",
      statusFailed: "Activation failed",
      activationRunning: "Activating...",
      credential: "Credential",
      disabled: "Disabled",
      credentialDisabled: "Enable the credential before trying to wake it up.",
      loadingCredentials: "Refreshing...",
      keyFailed: "Management key verification failed. Check it and try again.",
      keyPassed: "Management key verified.",
      managementKeyPlaceholder: "Enter CPA management key",
      missingCredential: "Select a credential first.",
      missingKey: "Enter the CPA management key first.",
      missingModel: "Select an enabled activation model first.",
      model: "Activation model",
      noAuthFiles: "No activatable credentials found.",
      noModels: "This credential has no enabled activation model group.",
      refresh: "Refresh credentials",
      statusBusy: "Busy",
      statusSkipped: "Skipped",
      statusSuccess: "Success",
      statusUnknown: "Unknown",
      title: "Quota activation",
      usable: "Available",
      verifyKey: "Verify key",
      verifyingKey: "Verifying...",
      manualTab: "Manual Activation",
      helpTab: "Help",
      searchPlaceholder: "Search credentials...",
      languageLabel: "Language",
      noMatchingCredential: "No matching credentials.",
      helpConfigTitle: "Configuration Fields",
      helpConfigDesc: "The fields required to configure the quota-activation plugin are as follows:",
      helpActivationModelsTitle: "Automatic activation model fields",
      helpActivationModelsDesc: "Enter model names directly; no JSON object is required:",
      helpAutoActivate: "Whether to enable automatic quota activation.",
      helpEnableBefore: "Whether to automatically enable disabled credentials before activation.",
      helpScanInterval: "Optional; automatic scan interval in minutes. Enter a number only (default 30).",
      helpActivationTimeout: "Optional; activation request timeout in seconds. Enter a number only (default 60).",
      helpMaxConcurrency: "Optional; concurrency limit for automatic scan activation (worker pool size); default 1.",
      helpActivationPrompt: "Optional; activation prompt text.",
      helpCodexModels: "Codex automatic activation model name, for example gpt-5-mini.",
      helpAntigravityGroup: "Antigravity automatic activation model group: gemini or claude_gpt.",
      helpAntigravityModels: "Model name for the current Antigravity model group.",
      recordsTab: "Execution Records",
      recordsTitle: "Execution Records",
      recordsHint: "Last 5 runs in process (cleared on restart)",
      refreshRecords: "Refresh records",
      noRecords: "No execution records",
      noRecordsAutoOff: "Auto activation is off. Set auto_activate to true in config; each scan tick will be recorded, and manual activation also writes records.",
      triggerManual: "Manual activation",
      triggerAuto: "Auto activation",
      kindScan: "Scheduled scan",
      kindActivationAuto: "Auto activation",
      kindActivationManual: "Manual activation",
      providerBreakdown: "By provider",
      statAttempted: "Attempted",
      statSucceeded: "Succeeded",
      statFailed: "Failed",
      statSkipped: "Skipped",
      loadingRecords: "Refreshing...",
      pageDesc: "quota-activation plugin: inspect credentials and trigger quota activation manually; execution records and configuration help live here.",
      pluginLabel: "Plugin",
      gateTitle: "Management key",
      gateDesc: "This resource page opens anonymously; inspecting credentials and running activation require the CPA management key, used only for this request.",
      managementKeyLabel: "CPA management key",
      manualNote: "Pick a credential and activation model, then trigger manually. You will confirm once more before execution, and cannot double-submit while running.",
      chipCredentials: "Credentials",
      chipAuto: "Auto activation",
      autoOn: "On",
      autoOff: "Off",
      autoUnknown: "—",
      confirmTitle: "Confirm activation",
      confirmCost: "Execution sends one model request and may consume the selected credential's model quota.",
      confirmRun: "Confirm and run",
      confirmCancel: "Cancel",
      confirmTargetCredentials: "Credentials",
      confirmTargetModel: "Model",
      recordsScopeNote: "Only in-process, already-redacted diagnostic summaries are shown."
    }
  };
  let language = "zh-CN";
  let authFiles = [];
  let activeProvider = "";
  let selectedModel = null;
  let runHistory = [];
  let autoActivateEnabled = null;
  let activationScope = null;
  // 显式共享 in-flight：并发/重复确认只允许一次执行，不依赖按钮 disabled 这种表现层状态。
  let activationInFlight = false;
  function hostLanguage() {
    const embedded = window.self !== window.top;
    if (embedded) {
      const hostLang = document.documentElement.dataset.hostLang || "";
      if (hostLang) { return hostLang.toLowerCase().indexOf("en") === 0 ? "en-US" : "zh-CN"; }
      return "zh-CN";
    }
    const nav = (navigator.language || "").toLowerCase();
    return nav.indexOf("en") === 0 ? "en-US" : "zh-CN";
  }
  function textFor(key) { return translations[language][key] || translations["zh-CN"][key] || key; }
  function message(text) { const toast = document.getElementById("managementToast"); toast.textContent = text || ""; }
  function managementKey() { const key = document.getElementById("managementKey").value.trim(); if (!key) { message(textFor("missingKey")); } return key; }
  function applyLanguage() {
    document.documentElement.lang = language;
    document.querySelectorAll("[data-i18n]").forEach((item) => { item.textContent = textFor(item.dataset.i18n); });
    document.querySelectorAll("[data-i18n-placeholder]").forEach((item) => { item.placeholder = textFor(item.dataset.i18nPlaceholder); });
    renderAuthFiles();
    renderRunHistory();
    updateHeaderChips();
  }
  function switchLanguage() { language = language === "zh-CN" ? "en-US" : "zh-CN"; applyLanguage(); message(""); }
  function shortError(defaultKey) { return textFor(defaultKey); }
  function setButtonBusy(control, labelKey) { if (!control) { return null; } const state = { control, labelKey: control.dataset.i18n || "", text: control.textContent }; control.disabled = true; control.textContent = textFor(labelKey); return state; }
  function restoreButton(state) { if (!state) { return; } state.control.disabled = false; state.control.textContent = state.labelKey ? textFor(state.labelKey) : state.text; }
  function activationErrorMessage(error) { const raw = String(error && error.message ? error.message : error); try { const payload = JSON.parse(raw); return payload.last_error || payload.message || raw || textFor("activationFailed"); } catch (_) { return raw || textFor("activationFailed"); } }
  async function managementFetch(path, options) {
    const key = managementKey();
    if (!key) { const error = new Error(textFor("missingKey")); error.code = "missing_key"; throw error; }
    const response = await fetch(path, { ...(options || {}), headers: { "Authorization": "Bearer " + key, "Content-Type": "application/json", ...((options && options.headers) || {}) } });
    const text = await response.text();
    if (!response.ok) { throw new Error(text || response.statusText); }
    return text ? JSON.parse(text) : {};
  }
  async function verifyManagementKey(button) {
    const key = managementKey();
    if (!key) { message(textFor("missingKey")); return; }
    const buttonState = setButtonBusy(button, "verifyingKey");
    try {
      await managementFetch(STATUS_PATH, { method: "GET" });
      document.getElementById("loginGate").hidden = true;
      document.getElementById("appShell").hidden = false;
      await loadAuthFiles();
      message(textFor("keyPassed"));
    } catch (error) {
      message(shortError("keyFailed"));
    } finally {
      restoreButton(buttonState);
    }
  }
  function setManualActionBusy(busy, labelKey) {
    const refreshBtn = document.getElementById("refreshCredentialsButton");
    const activateBtn = document.getElementById("manualActivateButton");
    const confirmRun = document.getElementById("confirmRunButton");
    const confirmCancel = document.getElementById("confirmCancelButton");
    const state = { refresh: null, activate: null, confirmRun: null, confirmCancel: null };
    if (busy) {
      if (refreshBtn) {
        state.refresh = { disabled: refreshBtn.disabled, text: refreshBtn.textContent, labelKey: refreshBtn.dataset.i18n || "refresh" };
        refreshBtn.disabled = true;
        if (labelKey) { refreshBtn.textContent = textFor(labelKey); }
      }
      if (activateBtn) {
        state.activate = { disabled: activateBtn.disabled, text: activateBtn.textContent, labelKey: activateBtn.dataset.i18n || "activate" };
        activateBtn.disabled = true;
      }
      if (confirmRun) { state.confirmRun = { disabled: confirmRun.disabled }; confirmRun.disabled = true; }
      if (confirmCancel) { state.confirmCancel = { disabled: confirmCancel.disabled }; confirmCancel.disabled = true; }
    }
    return state;
  }
  function restoreManualActionBusy(state) {
    if (!state) { return; }
    const refreshBtn = document.getElementById("refreshCredentialsButton");
    const activateBtn = document.getElementById("manualActivateButton");
    const confirmRun = document.getElementById("confirmRunButton");
    const confirmCancel = document.getElementById("confirmCancelButton");
    if (refreshBtn && state.refresh) { refreshBtn.disabled = Boolean(state.refresh.disabled); refreshBtn.textContent = state.refresh.labelKey ? textFor(state.refresh.labelKey) : state.refresh.text; }
    if (activateBtn && state.activate) { activateBtn.disabled = Boolean(state.activate.disabled); activateBtn.textContent = state.activate.labelKey ? textFor(state.activate.labelKey) : state.activate.text; }
    if (confirmRun && state.confirmRun) { confirmRun.disabled = Boolean(state.confirmRun.disabled); }
    if (confirmCancel && state.confirmCancel) { confirmCancel.disabled = Boolean(state.confirmCancel.disabled); }
  }
  async function loadAuthFiles(button) {
    void button;
    const manualBusy = setManualActionBusy(true, "loadingCredentials");
    try {
      const result = await managementFetch(AUTH_FILES_PATH, { method: "GET" });
      if (!result) { return; }
      authFiles = Array.isArray(result.files) ? result.files : [];
      await loadCredentialModels();
      renderAuthFiles();
      syncModelOptions();
      updateHeaderChips();
      if (authFiles.length === 0) { message(textFor("noAuthFiles")); }
    } catch (error) {
      message(activationErrorMessage(error));
    } finally {
      restoreManualActionBusy(manualBusy);
    }
  }
  async function loadCredentialModels() {
    const providerCache = {};
    const pendingByProvider = {};
    for (const credential of authFiles) {
      if (!credential) { continue; }
      if (Array.isArray(credential.models) && credential.models.length > 0) { continue; }
      const provider = String(credential.provider || "").toLowerCase();
      const cacheKey = provider || ("name:" + String(credential.name || credential.auth_id || ""));
      if (!credential.name && !providerCache[cacheKey]) { credential.models = Array.isArray(credential.models) ? credential.models : []; continue; }
      try {
        if (!Object.prototype.hasOwnProperty.call(providerCache, cacheKey)) {
          if (!pendingByProvider[cacheKey]) {
            const sampleName = credential.name || (authFiles.find((item) => item && item.name && String(item.provider || "").toLowerCase() === provider) || {}).name;
            pendingByProvider[cacheKey] = sampleName
              ? managementFetch(AUTH_FILE_MODELS_PATH + "?name=" + encodeURIComponent(sampleName), { method: "GET" }).then((result) => Array.isArray(result && result.models) ? result.models : []).catch(() => [])
              : Promise.resolve([]);
          }
          providerCache[cacheKey] = await pendingByProvider[cacheKey];
        }
        credential.models = modelChoicesForCredential(credential, providerCache[cacheKey] || []);
      } catch (_) {
        credential.models = Array.isArray(credential.models) ? credential.models : [];
      }
    }
  }
  function modelChoicesForCredential(credential, models) { const choices = []; for (const item of models) { const value = String((item && (item.id || item.name || item.value)) || "").trim(); if (!value) { continue; } if (credential.provider === "antigravity") { const group = antigravityModelGroup(value); if (!group) { continue; } choices.push({ value, label: (group === "gemini" ? "Gemini" : "Claude/GPT") + " · " + value, group }); } else { choices.push({ value, label: "Codex · " + value, group: "" }); } } return choices; }
  function antigravityModelGroup(model) { const lower = String(model || "").toLowerCase(); if (lower.includes("gemini")) { return "gemini"; } if (lower.includes("claude") || lower.includes("gpt")) { return "claude_gpt"; } return ""; }
  function setTabState(tab) {
    const mapping = [["manual", "manualTabButton"], ["records", "recordsTabButton"], ["help", "helpTabButton"]];
    for (const [name, id] of mapping) {
      const button = document.getElementById(id);
      const selected = name === tab;
      button.classList.toggle("active", selected);
      button.setAttribute("aria-selected", selected ? "true" : "false");
      button.tabIndex = selected ? 0 : -1;
      document.getElementById(name + "Panel").hidden = !selected;
    }
  }
  function switchTab(tab) {
    setTabState(tab);
    closeConfirm();
    if (tab === "records") { loadRunHistory(); }
  }
  function providerLabel(key) {
    if (key === "antigravity") return "Antigravity";
    if (key === "codex") return "Codex";
    return language === "zh-CN" ? "其他" : "Other";
  }
  function normalizeProviderKey(provider) {
    const key = String(provider || "").toLowerCase();
    if (key === "antigravity" || key === "codex") return key;
    return "other";
  }
  function credentialProvider(file) { return file ? file["provider"] : ""; }
  function availableProviders() {
    const seen = new Set();
    const order = ["antigravity", "codex", "other"];
    for (const file of authFiles) { seen.add(normalizeProviderKey(credentialProvider(file))); }
    return order.filter((key) => seen.has(key));
  }
  function switchProvider(provider) { activeProvider = provider || ""; renderAuthFiles(); syncModelOptions(); }
  function renderProviderTabs(providers) {
    const tabs = document.getElementById("providerTabs");
    if (!tabs) { return; }
    tabs.innerHTML = "";
    if (providers.length <= 1) { tabs.hidden = true; return; }
    tabs.hidden = false;
    for (const key of providers) {
      const button = document.createElement("button");
      button.type = "button";
      button.setAttribute("role", "radio");
      button.className = "provider-tab";
      button.setAttribute("aria-checked", key === activeProvider ? "true" : "false");
      button.tabIndex = key === activeProvider ? 0 : -1;
      button.dataset.provider = key;
      button.textContent = providerLabel(key);
      button.onclick = () => switchProvider(key);
      tabs.appendChild(button);
    }
  }
  function renderAuthFiles() {
    const list = document.getElementById("credentialSelect");
    if (!list) { return; }
    const providers = availableProviders();
    if (!activeProvider || !providers.includes(activeProvider)) { activeProvider = providers.length > 0 ? providers[0] : ""; }
    renderProviderTabs(providers);
    const searchVal = (document.getElementById("credentialSearch")?.value || "").toLowerCase().trim();
    const selectedValues = Array.from(list.querySelectorAll('input[type="checkbox"]:checked')).map((input) => input.value);
    list.innerHTML = "";
    const filtered = authFiles.filter(file => {
      if (normalizeProviderKey(credentialProvider(file)) !== activeProvider) return false;
      if (!searchVal) return true;
      const label = (file.label || "").toLowerCase();
      const authId = (file.auth_id || "").toLowerCase();
      return label.includes(searchVal) || authId.includes(searchVal);
    });
    for (const file of filtered) {
      const row = document.createElement("label");
      row.className = "credential-item";
      row.classList.toggle("disabled", Boolean(file.disabled));
      row.onclick = () => { if (!file.disabled) { return; } message(textFor("credentialDisabled")); };
      const checkbox = document.createElement("input");
      checkbox.type = "checkbox";
      checkbox.value = file.auth_id;
      checkbox.disabled = Boolean(file.disabled);
      checkbox.checked = !file.disabled && selectedValues.includes(file.auth_id);
      checkbox.onchange = () => syncModelOptions();
      const text = document.createElement("span");
      text.textContent = [file.label || file.auth_id, file.disabled ? textFor("disabled") : textFor("usable")].filter(Boolean).join(" · ");
      row.appendChild(checkbox);
      row.appendChild(text);
      list.appendChild(row);
    }
    if (filtered.length === 0) {
      const empty = document.createElement("div");
      empty.className = "empty-state";
      empty.textContent = authFiles.length === 0 ? textFor("noAuthFiles") : textFor("noMatchingCredential");
      list.appendChild(empty);
    }
  }
  function setModelMenuOpen(open) {
    const panel = document.getElementById("modelMenuPanel");
    const button = document.getElementById("modelMenuButton");
    if (!panel || !button) { return; }
    panel.classList.toggle("open", !!open);
    button.setAttribute("aria-expanded", open ? "true" : "false");
  }
  function toggleModelMenu() { const panel = document.getElementById("modelMenuPanel"); if (!panel) { return; } setModelMenuOpen(!panel.classList.contains("open")); }
  function chooseModel(value, group, label) {
    selectedModel = value ? { value, group: group || "", label: label || value } : null;
    const modelLabel = document.getElementById("modelMenuLabel");
    if (modelLabel) { modelLabel.textContent = selectedModel ? selectedModel.label : "—"; }
    const panel = document.getElementById("modelMenuPanel");
    if (panel) { panel.querySelectorAll(".model-option").forEach((btn) => { btn.classList.toggle("active", !!(selectedModel && btn.dataset.value === selectedModel.value)); }); }
    setModelMenuOpen(false);
  }
  function syncModelOptions() {
    const credentials = selectedCredentials();
    const panel = document.getElementById("modelMenuPanel");
    if (!panel) { return; }
    const selectedVal = selectedModel ? selectedModel.value : "";
    panel.innerHTML = "";
    const seen = new Set();
    const choices = [];
    for (const credential of credentials) {
      const models = credential && Array.isArray(credential.models) ? credential.models : [];
      for (const item of models) { if (!seen.has(item.value)) { seen.add(item.value); choices.push(item); } }
    }
    const selectedGroup = selectedModel ? selectedModel.group : "";
    for (const item of choices) {
      const option = document.createElement("button");
      option.type = "button";
      option.disabled = Boolean(selectedGroup && item.group && item.group !== selectedGroup);
      option.className = "model-option" + (item.value === selectedVal ? " active" : "") + (option.disabled ? " disabled" : "");
      option.dataset.value = item.value;
      option.dataset.group = item.group || "";
      option.textContent = item.label || item.value;
      option.onclick = () => { if (option.disabled) { return; } chooseModel(item.value, item.group || "", item.label || item.value); };
      panel.appendChild(option);
    }
    if (choices.length === 0) { chooseModel("", "", "—"); }
    else if (!choices.some((item) => item.value === selectedVal)) { chooseModel(choices[0].value, choices[0].group || "", choices[0].label || choices[0].value); }
    else { const current = choices.find((item) => item.value === selectedVal); chooseModel(current.value, current.group || "", current.label || current.value); }
    if (credentials.length > 0 && choices.length === 0) {
      const hasEnabled = credentials.some((credential) => credential && !credential.disabled);
      if (hasEnabled) { message(textFor("noModels")); }
    }
  }
  function selectedCredentials() {
    const list = document.getElementById("credentialSelect");
    if (!list) return [];
    const selectedValues = Array.from(list.querySelectorAll('input[type="checkbox"]:checked')).map((input) => input.value);
    return authFiles.filter((item) => selectedValues.includes(item.auth_id));
  }
  function selectedModelOption() { if (!selectedModel || !selectedModel.value) { return null; } return { value: selectedModel.value, dataset: { group: selectedModel.group || "" } }; }
  function quotaPayloadForActivation(credential) { return credential ? (credential.quota_payload || {}) : {}; }
  function updateHeaderChips() {
    const chips = document.getElementById("headerChips");
    if (!chips) { return; }
    chips.hidden = document.getElementById("appShell").hidden;
    const count = document.getElementById("chipCredCount");
    if (count) { count.textContent = String(authFiles.length); }
    const auto = document.getElementById("chipAuto");
    if (auto) { auto.textContent = autoActivateEnabled === true ? textFor("autoOn") : autoActivateEnabled === false ? textFor("autoOff") : textFor("autoUnknown"); }
  }
  function closeConfirm() {
    activationScope = null;
    const box = document.getElementById("confirmBox");
    if (box) { box.hidden = true; }
    const targets = document.getElementById("confirmTargets");
    if (targets) { targets.innerHTML = ""; }
  }
  function requestActivation(button) {
    const credentials = selectedCredentials();
    if (credentials.length === 0) { message(textFor("missingCredential")); return; }
    const modelOption = selectedModelOption();
    if (!modelOption) { message(textFor("missingModel")); return; }
    closeConfirm();
    activationScope = { credentials, modelOption, button: button || document.getElementById("manualActivateButton") };
    const targets = document.getElementById("confirmTargets");
    targets.innerHTML = "";
    const credentialItem = document.createElement("li");
    credentialItem.textContent = textFor("confirmTargetCredentials") + ": " + credentials.map((item) => item.label || item.auth_id).join(", ");
    const modelItem = document.createElement("li");
    modelItem.textContent = textFor("confirmTargetModel") + ": " + modelOption.value;
    targets.appendChild(credentialItem);
    targets.appendChild(modelItem);
    document.getElementById("confirmBox").hidden = false;
    document.getElementById("confirmRunButton").focus();
  }
  function cancelActivation() { closeConfirm(); message(""); }
  async function confirmActivation(button) {
    if (!activationScope) { return; }
    if (activationInFlight) { return; }
    activationInFlight = true;
    const credentials = activationScope.credentials;
    const modelOption = activationScope.modelOption;
    const actBtn = document.getElementById("manualActivateButton");
    const refBtn = document.getElementById("refreshCredentialsButton");
    const actState = setButtonBusy(actBtn, "activationRunning");
    const confirmRunState = setButtonBusy(button || document.getElementById("confirmRunButton"), "activationRunning");
    let refBtnDisabledState = false;
    if (refBtn) { refBtnDisabledState = refBtn.disabled; refBtn.disabled = true; }
    const cancelBtn = document.getElementById("confirmCancelButton");
    let cancelDisabledState = false;
    if (cancelBtn) { cancelDisabledState = cancelBtn.disabled; cancelBtn.disabled = true; }
    try {
      let successCount = 0;
      let failedCount = 0;
      let lastError = null;
      for (const credential of credentials) {
        const payload = {
          auth_id: credential.auth_id,
          provider: credential.provider,
          model_group: modelOption.dataset.group || "",
          model: modelOption.value,
          disabled: Boolean(credential.disabled),
          quota_payload: quotaPayloadForActivation(credential)
        };
        try {
          await managementFetch(ACTIVATE_PATH, { method: "POST", body: JSON.stringify(payload) });
          successCount++;
        } catch (error) {
          failedCount++;
          lastError = error;
        }
      }
      if (failedCount === 0) { message(textFor("activationDone")); }
      else if (lastError && (lastError.message === textFor("missingCredential") || lastError.message === textFor("missingModel") || lastError.message === textFor("missingKey"))) { message(lastError.message); }
      else { message(activationErrorMessage(lastError)); }
    } catch (error) {
      message(activationErrorMessage(error));
    } finally {
      activationInFlight = false;
      restoreButton(actState);
      restoreButton(confirmRunState);
      if (refBtn) { refBtn.disabled = refBtnDisabledState; }
      if (cancelBtn) { cancelBtn.disabled = cancelDisabledState; }
      closeConfirm();
    }
  }
  async function loadRunHistory(button) {
    const btnState = setButtonBusy(button, "loadingRecords");
    try {
      const result = await managementFetch(DIAGNOSTICS_PATH, { method: "GET" });
      if (!result) { return; }
      runHistory = Array.isArray(result.run_history) ? result.run_history : [];
      autoActivateEnabled = result.config && typeof result.config.auto_activate === "boolean" ? result.config.auto_activate : null;
      renderRunHistory();
      updateHeaderChips();
    } catch (error) {
      message(activationErrorMessage(error));
    } finally {
      restoreButton(btnState);
    }
  }
  function formatTime(isoStr) {
    if (!isoStr) return "—";
    try {
      const d = new Date(isoStr);
      if (isNaN(d.getTime())) return isoStr;
      return d.toLocaleString(language === "zh-CN" ? "zh-CN" : "en-US", { year: "numeric", month: "2-digit", day: "2-digit", hour: "2-digit", minute: "2-digit", second: "2-digit", hour12: false });
    } catch (_) { return isoStr; }
  }
  function summarySeparator() { return language === "zh-CN" ? "：" : ": "; }
  function kindLabel(item) {
    const kind = String(item && item.kind || "").toLowerCase();
    const trigger = String(item && item.trigger || "").toLowerCase();
    if (kind === "scan" || (kind === "activation" && trigger === "auto")) { return textFor("kindScan"); }
    if (trigger === "auto") { return textFor("kindActivationAuto"); }
    return textFor("kindActivationManual");
  }
  function kindBadgeClass(item) {
    const kind = String(item && item.kind || "").toLowerCase();
    const trigger = String(item && item.trigger || "").toLowerCase();
    if (kind === "scan" || (kind === "activation" && trigger === "auto")) { return "scan"; }
    return trigger === "auto" ? "auto" : "manual";
  }
  function localizeRecordMessage(raw) {
    const msg = String(raw || "").trim();
    if (!msg) { return ""; }
    if (language !== "zh-CN") {
      if (/状态保存被中断|状态保存失败/.test(msg)) { return "Activation succeeded, but saving state was interrupted (internal timeout/cancel); credential is still valid"; }
      if (msg.indexOf("调度器未选中目标凭证") >= 0) { return "Activation scheduler did not select target credential"; }
      if (/宿主模型执行器不可用/.test(msg) || /host model executor is unavailable/i.test(msg) || (/host_call_failed/i.test(msg) && /executor/i.test(msg))) { return "Host model executor unavailable (host not ready or callback failed; not a credential failure)"; }
      if (/定时扫描：/.test(msg)) {
        return msg.replace("定时扫描：", "Scheduled scan: ").replace(/尝试 /g, "attempted ").replace(/成功 /g, "succeeded ").replace(/失败 /g, "failed ").replace(/跳过 /g, "skipped ").replace(/；跳过：/, "; skip: ").replace(/；失败：/, "; fail: ").replace(/本周期已处理/g, "cycle already processed").replace(/执行中/g, "busy").replace(/宿主模型执行器不可用（已冷却）/g, "host model executor unavailable (cooled down)");
      }
      const enMap = [
        [/Codex唤醒失败：凭证已失效或鉴权失败（已重试3次）/, "Codex activation failed: credential invalid or unauthorized (retried 3 times)"],
        [/Antigravity唤醒失败：凭证已失效或鉴权失败（已重试3次）/, "Antigravity activation failed: credential invalid or unauthorized (retried 3 times)"],
        [/Codex唤醒失败：上游返回非成功状态（HTTP\s*(\d+)）/, "Codex activation failed: upstream returned a non-success status (HTTP $1)"],
        [/Antigravity唤醒失败：上游返回非成功状态（HTTP\s*(\d+)）/, "Antigravity activation failed: upstream returned a non-success status (HTTP $1)"],
        [/Codex唤醒失败：上游返回非成功状态/, "Codex activation failed: upstream returned a non-success status"],
        [/Antigravity唤醒失败：上游返回非成功状态/, "Antigravity activation failed: upstream returned a non-success status"],
        [/Codex唤醒失败：用量额度已耗尽/, "Codex activation failed: usage quota exhausted"],
        [/Codex唤醒失败：响应体为空/, "Codex activation failed: empty response body"],
        [/Antigravity唤醒失败：响应体为空/, "Antigravity activation failed: empty response body"],
        [/Codex唤醒失败：响应不是合法数据/, "Codex activation failed: response is not valid data"],
        [/Antigravity唤醒失败：响应不是合法数据/, "Antigravity activation failed: response is not valid data"],
        [/Codex唤醒失败：响应缺少有效输出结构/, "Codex activation failed: response missing valid output structure"],
        [/Codex唤醒失败：上游返回业务错误/, "Codex activation failed: upstream returned a business error"],
        [/Antigravity唤醒失败：上游返回业务错误/, "Antigravity activation failed: upstream returned a business error"],
        [/Codex唤醒失败：上游返回错误/, "Codex activation failed: upstream returned an error"],
        [/Antigravity唤醒失败：响应缺少选项/, "Antigravity activation failed: response missing choices"],
        [/Codex直连唤醒失败：无法读取凭证材料/, "Codex direct wake failed: unable to read credential material"],
        [/Antigravity直连唤醒失败：无法读取凭证材料/, "Antigravity direct wake failed: unable to read credential material"],
        [/本周期已唤醒，等待重置/, "Already woken this cycle; waiting for reset"],
        [/额度周期已处理/, "Quota cycle already processed"],
        [/读取物理凭证文件失败[：:]\s*(.*)/, "Failed to read physical credential file: $1"],
        [/读取物理凭证文件失败/, "Failed to read physical credential file"],
        [/读取凭证文件失败[：:]\s*(.*)/, "Failed to read credential file: $1"],
        [/读取凭证文件失败/, "Failed to read credential file"],
        [/直连唤醒失败[：:]\s*(.*)/, "Direct wake activation failed: $1"],
        [/直连唤醒失败/, "Direct wake activation failed"],
        [/列举凭证以解析可用模型失败[：:]\s*(.*)/, "Failed to list credentials for available models: $1"],
        [/列举凭证以解析可用模型失败/, "Failed to list credentials for available models"],
        [/列举凭证失败[：:]\s*(.*)/, "Failed to list credentials: $1"],
        [/列举凭证失败/, "Failed to list credentials"],
        [/缺少可用模型/, "Missing available models"],
        [/网络连接失败[：:]\s*(.*)/, "Network connection failed: $1"],
        [/网络连接失败/, "Network connection failed"],
        [/^执行中$/, "In progress"],
        [/唤醒请求无效[：:]\s*(.*)/, "Invalid activation request: $1"],
        [/未找到匹配的凭证文件/, "Matching credential file not found"],
        [/凭证已禁用/, "Credential disabled"],
        [/解析凭证文件失败/, "Failed to decode credential file"],
        [/写入凭证文件失败/, "Failed to write credential file"],
        [/宿主模型执行失败/, "Host model execution failed"],
        [/操作已取消/, "Operation canceled"],
        [/操作超时/, "Operation timed out"],
        [/提升凭证优先级失败/, "Failed to boost credential priority"],
        [/科德克斯唤醒失败：/, "Codex activation failed: "],
        [/反重力唤醒失败：/, "Antigravity activation failed: "]
      ];
      for (const pair of enMap) { if (pair[0].test(msg)) { return msg.replace(pair[0], pair[1]); } }
      return msg;
    }
    const lower = msg.toLowerCase();
    if (lower.indexOf("save activation state") >= 0 || lower.indexOf("save state context") >= 0 || (lower.indexOf("context canceled") >= 0 || lower.indexOf("context cancelled") >= 0)) { return "唤醒已成功，但状态保存被中断（内部超时/取消），非凭证失效"; }
    if (lower.indexOf("activation scheduler did not select target credential") >= 0) { return "调度器未选中目标凭证"; }
    if (lower.indexOf("host model executor is unavailable") >= 0 || (lower.indexOf("host_call_failed") >= 0 && lower.indexOf("executor") >= 0) || (lower.indexOf("host callback host.model.execute") >= 0 && lower.indexOf("unavailable") >= 0)) { return "宿主模型执行器不可用（宿主未就绪或回调失败，非凭证失效）"; }
    if (lower.indexOf("host model execute") >= 0 && lower.indexOf("unavailable") >= 0) { return "宿主模型执行器不可用（宿主未就绪或回调失败，非凭证失效）"; }
    if (/auto scan tick attempted=/.test(msg)) {
      const m = msg.match(/attempted=(\d+)\s+succeeded=(\d+)\s+failed=(\d+)\s+skipped=(\d+)/i);
      let out = m ? ("定时扫描：尝试 " + m[1] + " · 成功 " + m[2] + " · 失败 " + m[3] + " · 跳过 " + m[4]) : msg;
      const skipPart = msg.split(/;\s*skip:\s*/i)[1];
      if (skipPart) { out += "；跳过：" + skipPart.replace(/quota 周期已处理/g, "本周期已处理").replace(/busy/gi, "执行中").replace(/×/g, " ×"); }
      return out;
    }
    if (/^skip:\s*/i.test(msg)) { return msg.replace(/^skip:\s*/i, "跳过：").replace(/quota 周期已处理/g, "本周期已处理"); }
    return msg;
  }
  function isWarningMessage(msg) {
    const text = String(msg || "");
    if (!text) { return false; }
    const lower = text.toLowerCase();
    return /状态保存|非凭证失效/.test(text) || lower.indexOf("save activation state") >= 0 || lower.indexOf("context canceled") >= 0 || lower.indexOf("context cancelled") >= 0 || lower.indexOf("interrupted") >= 0;
  }
  function recordTone(item) {
    const failed = Number(item && item.failed || 0);
    const succeeded = Number(item && item.succeeded || 0);
    const msg = localizeRecordMessage(item && item.message);
    if (failed > 0 && succeeded === 0) { return "fail"; }
    if (succeeded > 0 && (isWarningMessage(msg) || isWarningMessage(item && item.message))) { return "warn"; }
    if (failed > 0 && succeeded > 0) { return "warn"; }
    return "ok";
  }
  function providerDisplayName(name) {
    const raw = String(name || "").toLowerCase();
    if (raw === "antigravity") { return "Antigravity"; }
    if (raw === "codex") { return "Codex"; }
    if (raw === "xai" || raw === "x-ai") { return "xAI"; }
    return name || providerLabel("other");
  }
  function statPill(label, value) {
    const span = document.createElement("span");
    span.className = "record-stat";
    span.textContent = label + summarySeparator() + String(value == null ? 0 : value);
    return span;
  }
  function renderRunHistory() {
    const list = document.getElementById("recordsList");
    if (!list) { return; }
    list.innerHTML = "";
    if (!Array.isArray(runHistory) || runHistory.length === 0) {
      const empty = document.createElement("div");
      empty.className = "empty-state";
      empty.textContent = autoActivateEnabled === false ? textFor("noRecordsAutoOff") : textFor("noRecords");
      list.appendChild(empty);
      return;
    }
    for (const item of runHistory) {
      const tone = recordTone(item);
      const card = document.createElement("div");
      card.className = "record-card" + (tone === "ok" ? "" : (" " + tone));
      const head = document.createElement("div");
      head.className = "record-head";
      const meta = document.createElement("div");
      meta.className = "record-meta";
      const badge = document.createElement("span");
      badge.className = "record-badge " + kindBadgeClass(item);
      badge.textContent = kindLabel(item);
      const timeSpan = document.createElement("span");
      timeSpan.className = "record-time";
      timeSpan.textContent = formatTime(item.at);
      meta.appendChild(badge);
      meta.appendChild(timeSpan);
      head.appendChild(meta);
      card.appendChild(head);
      const stats = document.createElement("div");
      stats.className = "record-stats";
      stats.append(
        statPill(textFor("statSucceeded"), item.succeeded || 0),
        statPill(textFor("statFailed"), item.failed || 0),
        statPill(textFor("statSkipped"), item.skipped || 0)
      );
      card.appendChild(stats);
      if (Array.isArray(item.providers) && item.providers.length > 0) {
        const title = document.createElement("div");
        title.className = "record-time";
        title.textContent = textFor("providerBreakdown");
        card.appendChild(title);
        const provDiv = document.createElement("div");
        provDiv.className = "record-providers";
        for (const p of item.providers) {
          const row = document.createElement("div");
          row.className = "record-provider";
          const name = document.createElement("strong");
          name.textContent = providerDisplayName(p.name);
          row.appendChild(name);
          const line = document.createElement("span");
          line.textContent = [
            textFor("statSucceeded") + summarySeparator() + (p.succeeded || 0),
            textFor("statFailed") + summarySeparator() + (p.failed || 0),
            textFor("statSkipped") + summarySeparator() + (p.skipped || 0)
          ].join(" · ");
          row.appendChild(line);
          if (p.error) {
            const err = document.createElement("span");
            const mapped = localizeRecordMessage(p.error);
            err.className = isWarningMessage(mapped) || isWarningMessage(p.error) ? "prov-warn" : "prov-err";
            err.textContent = mapped;
            row.appendChild(err);
          }
          provDiv.appendChild(row);
        }
        card.appendChild(provDiv);
      }
      list.appendChild(card);
    }
  }
  document.addEventListener("click", (event) => {
    const menu = document.getElementById("modelSelect");
    if (menu && !menu.contains(event.target)) { setModelMenuOpen(false); }
  });
  const tabButtons = [document.getElementById("manualTabButton"), document.getElementById("recordsTabButton"), document.getElementById("helpTabButton")];
  tabButtons.forEach((button, index) => {
    button.addEventListener("keydown", (event) => {
      const target = event.key === "ArrowRight" ? (index + 1) % tabButtons.length
        : event.key === "ArrowLeft" ? (index + tabButtons.length - 1) % tabButtons.length
        : event.key === "Home" ? 0
        : event.key === "End" ? tabButtons.length - 1 : -1;
      if (target >= 0) { event.preventDefault(); switchTab(tabButtons[target].dataset.tab); tabButtons[target].focus(); }
    });
  });
  const providerTabsEl = document.getElementById("providerTabs");
  if (providerTabsEl) {
    providerTabsEl.addEventListener("keydown", (event) => {
      const buttons = Array.prototype.slice.call(providerTabsEl.querySelectorAll(".provider-tab"));
      if (buttons.length === 0) { return; }
      let index = buttons.indexOf(document.activeElement);
      if (index < 0) { index = 0; }
      const target = (event.key === "ArrowRight" || event.key === "ArrowDown") ? (index + 1) % buttons.length
        : (event.key === "ArrowLeft" || event.key === "ArrowUp") ? (index - 1 + buttons.length) % buttons.length
        : event.key === "Home" ? 0
        : event.key === "End" ? buttons.length - 1 : -1;
      if (target < 0) { return; }
      event.preventDefault();
      const provider = buttons[target].dataset.provider;
      switchProvider(provider);
      const refocused = providerTabsEl.querySelector('.provider-tab[data-provider="' + provider + '"]');
      if (refocused) { refocused.focus(); }
    });
  }
  language = hostLanguage();
  const languageButton = document.getElementById("languageMenuButton");
  if (languageButton) { languageButton.hidden = window.self !== window.top; }
  applyLanguage();
</script>
</body>
</html>`))
}
