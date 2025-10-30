// // #region DevTools 測試的前置條件
// // formContext 在 DevTools 是藉由 Xrm.Page 取得，Dataverse 則是藉由 executionContext 取得
// const formContext = Xrm.Page;
// // #endregion

// // #region 表單取值
// /**
//  * 取得目前 record ID、entityName、使用者資訊、環境 URL
//  */
// const id = formContext.data.entity.getId().replace(/[{}]/g, "");
// const entityName = formContext.data.entity.getEntityName();
// const user = Xrm.Utility.getGlobalContext().userSettings;
// const clientUrl = Xrm.Utility.getGlobalContext().getClientUrl();

// console.log({ id, entityName, userId: user.userId, userName: user.userName, clientUrl });
// // #endregion

// // #region 欄位操作（讀/寫/必填/顯示/唯讀）
// const attrName = "attrName";
// const attr = formContext.getAttribute(attrName);
// const ctrl = formContext.getControl(attrName);
// if (attr && ctrl) {
//   console.log(`目前欄位 ${attrName} 的值為 ${attr.getValue()}`);
//   attr.setValue("新的值"); // 設定欄位值
//   attr.setSubmitMode("always"); // 設定欄位提交模式 (預設為 "dirty")
//   attr.setRequiredLevel("required"); // 設定欄位必填狀態
//   ctrl.setVisible(true); // 設定欄位可見性
//   ctrl.setDisabled(true); // 設定欄位唯讀狀態
// }
// // #endregion

// // #region onChange 事件註冊
// const attrName = "attrName";
// const attr = formContext.getAttribute(attrName);
// // onChangeHandler 為業務邏輯函式

// // 1. 直接註冊 onChange 事件
// if (attr) {
//   attr.removeOnChange(onChangeHandler); // 移除 onChange 事件
//   attr.addOnChange(onChangeHandler); // 註冊 onChange 事件
// }
// // 2. 透過現有工具函式 addOnChangeEvent 註冊 onChange 事件
// addOnChangeEvent(formContext, attrName, onChangeHandler);
// // #endregion

// // #region 欄位/表單通知
// // 加入/移除 表單通知
// formContext.ui.setFormNotification("通知內容", "ERROR", "formNotificationId"); // 表單通知
// formContext.ui.clearFormNotification("formNotificationId"); // 移除表單通知

// // 加入/移除 欄位通知
// const attrName = "attrName";
// const ctrl = formContext.getControl(attrName);
// ctrl?.setNotification("通知內容", "ERROR", "attrNotificationId"); // 欄位通知
// ctrl?.clearNotification(); // 移除欄位所有通知
// // #endregion

// // #region OptionSet 操作
// const attrName = "attrName";
// const optionAttr = formContext.getAttribute(attrName);
// const optionCtrl = formContext.getControl(attrName);
// const optionVal = optionAttr?.getValue();
// const optionLbl = optionAttr?.getOption(optionVal)?.text ?? "找不到此 option 的 label";
// console.log("目前的 option:", { optionVal, optionLbl });

// // 設定 OptionSet 的值
// optionAttr?.setValue(1); // 需確認 dataverse 的 optionSet 值是否正確
// // #endregion

// // #region Lookup 欄位操作
// const attrName = "attrName";
// const lookupAttr = formContext.getAttribute(attrName);
// const lookupCtrl = formContext.getControl(attrName);
// const lookupVal = lookupAttr?.getValue(); // 取出來的 value 會是一個陣列
// const lookupEntityType = lookupVal[0]?.entityType ?? "找不到此 lookup 的 entityType";
// const lookupId = lookupVal[0]?.id ?? "找不到此 lookup 的 id";
// const lookupName = lookupVal[0]?.name ?? "找不到此 lookup 的 name";
// console.log("目前的 lookup:", { lookupEntityType, lookupId, lookupName });

// // 設定 Lookup 欄位的值
// lookupAttr?.setValue([{ id: "lookupId", name: "lookupName", entityType: "lookupEntityType" }]); // id 必須存在於 dataverse，name 可以自訂
// // #endregion

// // #region Subgrid 操作
// const gridCtrl = formContext.getControl("SubgridName"); // SubgridName 取自 subgrid 的 屬性 -> 名稱
// if (!gridCtrl) {
//   console.log("找不到指定的 subgrid 控制項，請檢查 subgrid 名稱是否正確！");
// }
// const grid = gridCtrl.getGrid();
// const selected = grid?.getSelectedRows();
// const ids = [];
// selected?.forEach(function (row) {
//   ids.push(row.getData().getEntity().getId().replace(/[{}]/g, ""));
// });
// console.log("勾選 IDs:", ids);

// // 刷新 subgrid
// gridCtrl?.refresh();
// // #endregion

// // #region 對話框/頁面導覽
// // 通知 dialog
// Xrm.Navigation.openAlertDialog({ text: "通知訊息" });
// // 確認 dialog
// Xrm.Navigation.openConfirmDialog({ text: "確定執行？" }).then((result) => {
//   console.log("confirmed:", result.confirmed);
// });
// // 錯誤 dialog
// Xrm.Navigation.openErrorDialog({
//   message: "發生錯誤",
//   details: "漢堡獵人快來救救我@@",
// });
// // 頁面導覽
// Xrm.Navigation.openForm({
//   entityName: "entityName",
//   entityId: "{recordId}",
// });
// // #endregion

// #region Xrm.WebApi CRUD 查詢

// #endregion
