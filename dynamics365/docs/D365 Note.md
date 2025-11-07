# WebResource

## JavaScript

### JS WebResource Library Template

- 應在專案 repo 的 .vscode/artsolutions.code-snippets 設定好，只需輸入 keyword: artLibrary 即可套用此模板

```code-snippets
{
  "ArtLibrary": {
    "scope": "javascript",
    "prefix": "artLibrary",
    "description": "JavaScript library template for Dynamics 365 Power Apps WebResource.",
    "body": [
      "\"use strict\";",
      "var art = art || {};",
      "",
      "art.${1:EntityName} = {",
      "  // #region setup functions",
      "  onLoad: function (executionContext) {",
      "    const formContext = executionContext.getFormContext();",
      "    this.setupForm(formContext);",
      "    this.setupVisible(formContext);",
      "    this.setupFilters(formContext);",
      "    this.setupEditable(formContext);",
      "    this.setupRequired(formContext);",
      "    this.setupOptions(formContext);",
      "    this.setupDefault(formContext);",
      "    this.setupOnChange(formContext);",
      "  },",
      "  onSave: function (executionContext) {",
      "    const formContext = executionContext.getFormContext();",
      "  },",
      "  setupForm: function (formContext) {},",
      "  setupVisible: function (formContext) {},",
      "  setupFilters: function (formContext) {},",
      "  setupEditable: function (formContext) {},",
      "  setupRequired: function (formContext) {},",
      "  setupOptions: function (formContext) {},",
      "  setupDefault: function (formContext) {},",
      "  setupOnChange: function (formContext) {",
      "    const { addOnChangeEvent } = art.${1:EntityName};",
      "",
      "    // addOnChangeEvent(formContext, \"fieldName\", art.${1:EntityName}.functionName);",
      "  },",
      "  // #endregion",
      "",
      "  // #region attribute functions",
      "  // art_column1_onchange: function (executionContext) {},",
      "  // art_column2_filter: function (formContext) {},",
      "  // art_column3_calculate: function (formContext) {},",
      "  // #endregion",
      "",
      "  // #region subgrid functions",
      "  // Subgrid1_onload: function (executionContext) {},",
      "  // Subgrid1_onsave: function (executionContext) {},",
      "  // Subgrid2_onselect: function (executionContext) {},",
      "  // #endregion",
      "",
      "  // #region local functions",
      "  /** 註冊欄位 `onChange` 事件",
      "   *",
      "   * @param {object} formContext - 表單上下文",
      "   * @param {string} attrName - 欄位邏輯名稱",
      "   * @param {function} eventHandler - 業務邏輯函式",
      "   */",
      "  addOnChangeEvent(formContext, attrName, eventHandler) {",
      "    const attr = formContext.getAttribute(attrName);",
      "    if (attr) {",
      "      attr.removeOnChange(eventHandler);",
      "      attr.addOnChange(eventHandler);",
      "    } else {",
      "      console.warn(`註冊 onChange 事件失敗：找不到欄位 ${attrName}`);",
      "    }",
      "  },",
      "  /** 檢查表單上的是否存在指定欄位",
      "   *",
      "   * @param {object} formContext - 表單上下文",
      "   * @param {string[]} fieldsToValidate - 要檢查的欄位名稱陣列",
      "   * @returns {boolean} 如果欄位都存在則回傳 true，否則回傳 false",
      "   */",
      "  validateFields: function (formContext, fieldsToValidate) {",
      "    for (let field of fieldsToValidate) {",
      "      const attr = formContext.getAttribute(field);",
      "      if (!attr) {",
      "        console.error(`validateFields: 欄位 '${field}' 不存在`);",
      "        return false;",
      "      }",
      "    }",
      "    return true;",
      "  },",
      "  // #endregion",
      "",
      "  // #region button functions",
      "  // ButtonName1_action: function (formContext) {},",
      "  // ButtonName1_enable: function (formContext) {",
      "  //   return false;",
      "  // },",
      "  // ButtonName2_action: function (formContext) {},",
      "  // ButtonName2_enable: function (formContext) {",
      "  //   return false;",
      "  // },",
      "  // #endregion",
      "};",
      ""
    ]
  }
}
```

```javascript
"use strict";
var art = art || {};

art.EntityName = {
  // #region setup functions
  onLoad: function (executionContext) {
    const formContext = executionContext.getFormContext();
    this.setupForm(formContext);
    this.setupVisible(formContext);
    this.setupFilters(formContext);
    this.setupEditable(formContext);
    this.setupRequired(formContext);
    this.setupOptions(formContext);
    this.setupDefault(formContext);
    this.setupOnChange(formContext);
  },
  onSave: function (executionContext) {
    const formContext = executionContext.getFormContext();
  },
  setupForm: function (formContext) {},
  setupVisible: function (formContext) {},
  setupFilters: function (formContext) {},
  setupEditable: function (formContext) {},
  setupRequired: function (formContext) {},
  setupOptions: function (formContext) {},
  setupDefault: function (formContext) {},
  setupOnChange: function (formContext) {
    const { addOnChangeEvent } = art.EntityName;

    // addOnChangeEvent(formContext, "fieldName", art.EntityName.functionName);
  },
  // #endregion

  // #region attribute functions
  // art_column1_onchange: function (executionContext) {},
  // art_column2_filter: function (formContext) {},
  // art_column3_calculate: function (formContext) {},
  // #endregion

  // #region subgrid functions
  // Subgrid1_onload: function (executionContext) {},
  // Subgrid1_onsave: function (executionContext) {},
  // Subgrid2_onselect: function (executionContext) {},
  // #endregion

  // #region local functions
  /** 註冊欄位 `onChange` 事件
   *
   * @param {object} formContext - 表單上下文
   * @param {string} attrName - 欄位邏輯名稱
   * @param {function} eventHandler - 業務邏輯函式
   */
  addOnChangeEvent(formContext, attrName, eventHandler) {
    const attr = formContext.getAttribute(attrName);
    if (attr) {
      attr.removeOnChange(eventHandler);
      attr.addOnChange(eventHandler);
    } else {
      console.warn(`註冊 onChange 事件失敗：找不到欄位 attrName`);
    }
  },
  /** 檢查表單上的是否存在指定欄位
   *
   * @param {object} formContext - 表單上下文
   * @param {string[]} fieldsToValidate - 要檢查的欄位名稱陣列
   * @returns {boolean} 如果欄位都存在則回傳 true，否則回傳 false
   */
  validateFields: function (formContext, fieldsToValidate) {
    for (let field of fieldsToValidate) {
      const attr = formContext.getAttribute(field);
      if (!attr) {
        console.error(`validateFields: 欄位 'field' 不存在`);
        return false;
      }
    }
    return true;
  },
  // #endregion

  // #region button functions
  // ButtonName1_action: function (formContext) {},
  // ButtonName1_enable: function (formContext) {
  //   return false;
  // },
  // ButtonName2_action: function (formContext) {},
  // ButtonName2_enable: function (formContext) {
  //   return false;
  // },
  // #endregion
};
```

### 常用語法

#### DevTools 測試的前置條件

- 在表單頁面開 DevTools (F12)

```javascript
// formContext 在 DevTools 是藉由 Xrm.Page 取得，Dataverse 則是藉由 executionContext 取得
const formContext = Xrm.Page;
```

#### 表單取值

```javascript
/**
 * 取得目前 record ID、entityName、使用者資訊、環境 URL
 */
const id = formContext.data.entity.getId().replace(/[{}]/g, "");
const entityName = formContext.data.entity.getEntityName();
const user = Xrm.Utility.getGlobalContext().userSettings;
const clientUrl = Xrm.Utility.getGlobalContext().getClientUrl();

console.log({ id, entityName, userId: user.userId, userName: user.userName, clientUrl });
```

---

### 父表單取值

```javascript
/**
 * 取得父表單 record 的 ID
 */
const params = new URLSearchParams(window.parent.location.search);
const parentRecordId = params.get('id');

console.log({ ParentRecordId: parentRecordId });
```

---
### 欄位操作（讀/寫/必填/顯示/唯讀）

```javascript
const attrName = "attrName";
const attr = formContext.getAttribute(attrName);
const ctrl = formContext.getControl(attrName);
if (attr && ctrl) {
  console.log(`目前欄位 ${attrName} 的值為 ${attr.getValue()}`);
  attr.setValue("新的值"); // 設定欄位值
  attr.setSubmitMode("always"); // 設定欄位提交模式 (預設為 "dirty")
  attr.setRequiredLevel("required"); // 設定欄位必填狀態
  ctrl.setVisible(true); // 設定欄位可見性
  ctrl.setDisabled(true); // 設定欄位唯讀狀態
}
```

#### onChange 事件註冊

```javascript
const attrName = "attrName";
const attr = formContext.getAttribute(attrName);
// onChangeHandler 為業務邏輯函式

// 1. 直接註冊 onChange 事件
if (attr) {
  attr.removeOnChange(onChangeHandler); // 移除 onChange 事件
  attr.addOnChange(onChangeHandler); // 註冊 onChange 事件
}
// 2. 透過現有工具函式 addOnChangeEvent 註冊 onChange 事件
addOnChangeEvent(formContext, attrName, onChangeHandler);
```

#### 欄位/表單通知

```javascript
// 加入/移除 表單通知
formContext.ui.setFormNotification("通知內容", "ERROR", "formNotificationId"); // 表單通知
formContext.ui.clearFormNotification("formNotificationId"); // 移除表單通知

// 加入/移除 欄位通知
const attrName = "attrName";
const ctrl = formContext.getControl(attrName);
ctrl?.setNotification("通知內容", "ERROR", "attrNotificationId"); // 欄位通知
ctrl?.clearNotification(); // 移除欄位所有通知
```

#### OptionSet 操作

```javascript
const attrName = "attrName";
const optionAttr = formContext.getAttribute(attrName);
const optionCtrl = formContext.getControl(attrName);
const optionVal = optionAttr?.getValue();
const optionLbl = optionAttr?.getOption(optionVal)?.text ?? "找不到此 option 的 label";
console.log("目前的 option:", { optionVal, optionLbl });

// 設定 OptionSet 的值
optionAttr?.setValue(1); // 需確認 dataverse 的 optionSet 值是否正確
```

#### Lookup 欄位操作

```javascript
const attrName = "attrName";
const lookupAttr = formContext.getAttribute(attrName);
const lookupCtrl = formContext.getControl(attrName);
const lookupVal = lookupAttr?.getValue(); // 取出來的 value 會是一個陣列
const lookupEntityType = lookupVal[0]?.entityType ?? "找不到此 lookup 的 entityType";
const lookupId = lookupVal[0]?.id ?? "找不到此 lookup 的 id";
const lookupName = lookupVal[0]?.name ?? "找不到此 lookup 的 name";
console.log("目前的 lookup:", { lookupEntityType, lookupId, lookupName });

// 設定 Lookup 欄位的值
lookupAttr?.setValue([{ id: "lookupId", name: "lookupName", entityType: "lookupEntityType" }]); // id 必須存在於 dataverse，name 可以自訂
```

#### Subgrid 操作

```javascript
const gridCtrl = formContext.getControl("SubgridName"); // SubgridName 取自 subgrid 的 屬性 -> 名稱
if (!gridCtrl) {
  console.log("找不到指定的 subgrid 控制項，請檢查 subgrid 名稱是否正確！");
}
const grid = gridCtrl.getGrid();
const selected = grid?.getSelectedRows();
const ids = [];
selected?.forEach(function (row) {
  ids.push(row.getData().getEntity().getId().replace(/[{}]/g, ""));
});
console.log("勾選 IDs:", ids);

// 刷新 subgrid
gridCtrl?.refresh();
```

#### 對話框/頁面導覽

```javascript
// 通知 dialog
Xrm.Navigation.openAlertDialog({ text: "通知訊息" });
// 確認 dialog
Xrm.Navigation.openConfirmDialog({ text: "確定執行？" }).then((result) => {
  console.log("confirmed:", result.confirmed);
});
// 錯誤 dialog
Xrm.Navigation.openErrorDialog({
  message: "發生錯誤",
  details: "漢堡獵人快來救救我@@",
});
// 頁面導覽
Xrm.Navigation.openForm({
  entityName: "entityName",
  entityId: "{recordId}",
});
```

#### Xrm.WebApi CRUD 查詢

```javascript
// 取單筆
Xrm.WebApi.retrieveRecord("entityName", "{recordId}", "?$select=attrName1,attrName2").then((res) => {
  console.log("entityName:", res);
});

// 取單筆（含 lookup 展開）
Xrm.WebApi.retrieveRecord(
  "entityName",
  "{recordId}",
  "?$select=attrName1,attrName2&$expand=lookupAttrName($select=lookupEntityAttrName1,lookupEntityAttrName2)"
).then((res) => {
  console.log("entityName:", res);
});

// 查多筆
Xrm.WebApi.retrieveMultipleRecords(
  "entityName",
  "?$select=attrName1, attrName2&$filter=attrName3 eq 'value'&$orderby=createdon desc&$top=10"
).then((res) => {
  console.table(res.entities);
});

// 建立
Xrm.WebApi.createRecord("entityName", {
  attrName: "testValue1",
}).then((res) => console.log("created id:", res.id));

// 更新
Xrm.WebApi.updateRecord("entityName", "{recordId}", {
  attrName: "testValue2",
}).then((res) => console.log("updated:", res));

// 刪除
Xrm.WebApi.deleteRecord("entityName", "{recordId}").then(() => console.log("deleted"));
```

#### 儲存前阻擋/驗證

```javascript
formContext.data.entity.addOnSave(function (executionContext) {
  const args = executionContext.getEventArgs();
  const attrVal = formContext.getAttribute("attrName")?.getValue();
  const errs = [];
  if (!attrVal) errs.push("欄位沒有值或找不到欄位");
  if (errs.length) {
    args.preventDefault(); // 阻擋儲存
  }
});
```

### 通用函數

#### 註冊欄位 `onChange` 事件 - addOnChangeEvent

```javascript
/** 註冊欄位 `onChange` 事件
   *
   * @param {object} formContext - 表單上下文
   * @param {string} attrName - 欄位邏輯名稱
   * @param {function} eventHandler - 業務邏輯函式
   */
  addOnChangeEvent(formContext, attrName, eventHandler) {
    const attr = formContext.getAttribute(attrName);
    if (attr) {
      attr.removeOnChange(eventHandler);
      attr.addOnChange(eventHandler);
    } else {
      console.warn(`註冊 onChange 事件失敗：找不到欄位 attrName`);
    }
  },
```

#### 檢查表單上的是否存在指定欄位 - validateFields

```javascript
/** 檢查表單上的是否存在指定欄位
   *
   * @param {object} formContext - 表單上下文
   * @param {string[]} fieldsToValidate - 要檢查的欄位名稱陣列
   * @returns {boolean} 如果欄位都存在則回傳 true，否則回傳 false
   */
  validateFields: function (formContext, fieldsToValidate) {
    for (let field of fieldsToValidate) {
      const attr = formContext.getAttribute(field);
      if (!attr) {
        console.error(`validateFields: 欄位 'field' 不存在`);
        return false;
      }
    }
    return true;
  },
```
