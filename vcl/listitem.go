
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TListItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewListItem(AOwner *TListItems) *TListItem {
    l := new(TListItem)
    l.instance = ListItem_Create(CheckPtr(AOwner))
    l.ptr = unsafe.Pointer(l.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(l, (*TListItem).Free)
    return l
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsListItem(obj interface{}) *TListItem {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TListItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsListItem.
func ListItemFromInst(inst uintptr) *TListItem {
    return AsListItem(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsListItem.
func ListItemFromObj(obj IObject) *TListItem {
    return AsListItem(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsListItem.
func ListItemFromUnsafePointer(ptr unsafe.Pointer) *TListItem {
    return AsListItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (l *TListItem) Free() {
    if l.instance != 0 {
        ListItem_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TListItem) Instance() uintptr {
    return l.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TListItem) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TListItem) IsValid() bool {
    return l.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TListItem) Is() TIs {
    return TIs(l.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TListItem) As() TAs {
//    return TAs(l.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TListItemClass() TClass {
    return ListItem_StaticClassType()
}

func (l *TListItem) DisplayRectSubItem(subItem int32, Code TDisplayCode) TRect {
    return ListItem_DisplayRectSubItem(l.instance, subItem , Code)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TListItem) Assign(Source IObject) {
    ListItem_Assign(l.instance, CheckPtr(Source))
}

func (l *TListItem) Delete() {
    ListItem_Delete(l.instance)
}

func (l *TListItem) DisplayRect(Code TDisplayCode) TRect {
    return ListItem_DisplayRect(l.instance, Code)
}

func (l *TListItem) EditCaption() bool {
    return ListItem_EditCaption(l.instance)
}

func (l *TListItem) MakeVisible(PartialOK bool) {
    ListItem_MakeVisible(l.instance, PartialOK)
}

// 获取类名路径。
//
// Get the class name path.
func (l *TListItem) GetNamePath() string {
    return ListItem_GetNamePath(l.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (l *TListItem) ClassType() TClass {
    return ListItem_ClassType(l.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TListItem) ClassName() string {
    return ListItem_ClassName(l.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TListItem) InstanceSize() int32 {
    return ListItem_InstanceSize(l.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TListItem) InheritsFrom(AClass TClass) bool {
    return ListItem_InheritsFrom(l.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (l *TListItem) Equals(Obj IObject) bool {
    return ListItem_Equals(l.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TListItem) GetHashCode() int32 {
    return ListItem_GetHashCode(l.instance)
}

// 文本类信息。
//
// Text information.
func (l *TListItem) ToString() string {
    return ListItem_ToString(l.instance)
}

func (l *TListItem) DropTarget() bool {
    return ListItem_GetDropTarget(l.instance)
}

func (l *TListItem) SetDropTarget(value bool) {
    ListItem_SetDropTarget(l.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (l *TListItem) Caption() string {
    return ListItem_GetCaption(l.instance)
}

// 设置控件标题。
//
// Set the control title.
func (l *TListItem) SetCaption(value string) {
    ListItem_SetCaption(l.instance, value)
}

// 获取是否选中。
func (l *TListItem) Checked() bool {
    return ListItem_GetChecked(l.instance)
}

// 设置是否选中。
func (l *TListItem) SetChecked(value bool) {
    ListItem_SetChecked(l.instance, value)
}

func (l *TListItem) Cut() bool {
    return ListItem_GetCut(l.instance)
}

func (l *TListItem) SetCut(value bool) {
    ListItem_SetCut(l.instance, value)
}

func (l *TListItem) Data() unsafe.Pointer {
    return ListItem_GetData(l.instance)
}

func (l *TListItem) SetData(value unsafe.Pointer) {
    ListItem_SetData(l.instance, value)
}

// 获取返回是否获取焦点。
//
// Get Return to get focus.
func (l *TListItem) Focused() bool {
    return ListItem_GetFocused(l.instance)
}

// 设置返回是否获取焦点。
//
// Set Return to get focus.
func (l *TListItem) SetFocused(value bool) {
    ListItem_SetFocused(l.instance, value)
}

// 获取图像在images中的索引。
func (l *TListItem) ImageIndex() int32 {
    return ListItem_GetImageIndex(l.instance)
}

// 设置图像在images中的索引。
func (l *TListItem) SetImageIndex(value int32) {
    ListItem_SetImageIndex(l.instance, value)
}

func (l *TListItem) Index() int32 {
    return ListItem_GetIndex(l.instance)
}

// 获取左边位置。
//
// Get Left position.
func (l *TListItem) Left() int32 {
    return ListItem_GetLeft(l.instance)
}

// 设置左边位置。
//
// Set Left position.
func (l *TListItem) SetLeft(value int32) {
    ListItem_SetLeft(l.instance, value)
}

func (l *TListItem) ListView() *TWinControl {
    return AsWinControl(ListItem_GetListView(l.instance))
}

// 获取组件所有者。
//
// Get component owner.
func (l *TListItem) Owner() *TListItems {
    return AsListItems(ListItem_GetOwner(l.instance))
}

func (l *TListItem) Position() TPoint {
    return ListItem_GetPosition(l.instance)
}

func (l *TListItem) SetPosition(value TPoint) {
    ListItem_SetPosition(l.instance, value)
}

func (l *TListItem) Selected() bool {
    return ListItem_GetSelected(l.instance)
}

func (l *TListItem) SetSelected(value bool) {
    ListItem_SetSelected(l.instance, value)
}

func (l *TListItem) StateIndex() int32 {
    return ListItem_GetStateIndex(l.instance)
}

func (l *TListItem) SetStateIndex(value int32) {
    ListItem_SetStateIndex(l.instance, value)
}

func (l *TListItem) SubItems() *TStrings {
    return AsStrings(ListItem_GetSubItems(l.instance))
}

func (l *TListItem) SetSubItems(value IObject) {
    ListItem_SetSubItems(l.instance, CheckPtr(value))
}

// 获取顶边位置。
//
// Get Top position.
func (l *TListItem) Top() int32 {
    return ListItem_GetTop(l.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (l *TListItem) SetTop(value int32) {
    ListItem_SetTop(l.instance, value)
}

func (l *TListItem) SubItemImages(Index int32) int32 {
    return ListItem_GetSubItemImages(l.instance, Index)
}

func (l *TListItem) SetSubItemImages(Index int32, value int32) {
    ListItem_SetSubItemImages(l.instance, Index, value)
}

