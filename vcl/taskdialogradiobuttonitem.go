
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
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

type TTaskDialogRadioButtonItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskDialogRadioButtonItem() *TTaskDialogRadioButtonItem {
    t := new(TTaskDialogRadioButtonItem)
    t.instance = TaskDialogRadioButtonItem_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTaskDialogRadioButtonItem(obj interface{}) *TTaskDialogRadioButtonItem {
    t := new(TTaskDialogRadioButtonItem)
    t.instance, t.ptr = getInstance(obj)
    if t.instance == 0 { return nil }
    return t
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskDialogRadioButtonItem.
func TaskDialogRadioButtonItemFromInst(inst uintptr) *TTaskDialogRadioButtonItem {
    return AsTaskDialogRadioButtonItem(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTaskDialogRadioButtonItem.
func TaskDialogRadioButtonItemFromObj(obj IObject) *TTaskDialogRadioButtonItem {
    return AsTaskDialogRadioButtonItem(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskDialogRadioButtonItem.
func TaskDialogRadioButtonItemFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogRadioButtonItem {
    return AsTaskDialogRadioButtonItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TTaskDialogRadioButtonItem) Free() {
    if t.instance != 0 {
        TaskDialogRadioButtonItem_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialogRadioButtonItem) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialogRadioButtonItem) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialogRadioButtonItem) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTaskDialogRadioButtonItem) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTaskDialogRadioButtonItem) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogRadioButtonItemClass() TClass {
    return TaskDialogRadioButtonItem_StaticClassType()
}

// CN: 单击。
// EN: .
func (t *TTaskDialogRadioButtonItem) Click() {
    TaskDialogRadioButtonItem_Click(t.instance)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialogRadioButtonItem) GetNamePath() string {
    return TaskDialogRadioButtonItem_GetNamePath(t.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialogRadioButtonItem) Assign(Source IObject) {
    TaskDialogRadioButtonItem_Assign(t.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskDialogRadioButtonItem) DisposeOf() {
    TaskDialogRadioButtonItem_DisposeOf(t.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialogRadioButtonItem) ClassType() TClass {
    return TaskDialogRadioButtonItem_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialogRadioButtonItem) ClassName() string {
    return TaskDialogRadioButtonItem_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialogRadioButtonItem) InstanceSize() int32 {
    return TaskDialogRadioButtonItem_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialogRadioButtonItem) InheritsFrom(AClass TClass) bool {
    return TaskDialogRadioButtonItem_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialogRadioButtonItem) Equals(Obj IObject) bool {
    return TaskDialogRadioButtonItem_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialogRadioButtonItem) GetHashCode() int32 {
    return TaskDialogRadioButtonItem_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialogRadioButtonItem) ToString() string {
    return TaskDialogRadioButtonItem_ToString(t.instance)
}

// CN: 获取模态对话框显示结果。
// EN: .
func (t *TTaskDialogRadioButtonItem) ModalResult() TModalResult {
    return TaskDialogRadioButtonItem_GetModalResult(t.instance)
}

// CN: 设置模态对话框显示结果。
// EN: .
func (t *TTaskDialogRadioButtonItem) SetModalResult(value TModalResult) {
    TaskDialogRadioButtonItem_SetModalResult(t.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (t *TTaskDialogRadioButtonItem) Caption() string {
    return TaskDialogRadioButtonItem_GetCaption(t.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (t *TTaskDialogRadioButtonItem) SetCaption(value string) {
    TaskDialogRadioButtonItem_SetCaption(t.instance, value)
}

func (t *TTaskDialogRadioButtonItem) Default() bool {
    return TaskDialogRadioButtonItem_GetDefault(t.instance)
}

func (t *TTaskDialogRadioButtonItem) SetDefault(value bool) {
    TaskDialogRadioButtonItem_SetDefault(t.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTaskDialogRadioButtonItem) Enabled() bool {
    return TaskDialogRadioButtonItem_GetEnabled(t.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTaskDialogRadioButtonItem) SetEnabled(value bool) {
    TaskDialogRadioButtonItem_SetEnabled(t.instance, value)
}

func (t *TTaskDialogRadioButtonItem) Collection() *TCollection {
    return AsCollection(TaskDialogRadioButtonItem_GetCollection(t.instance))
}

func (t *TTaskDialogRadioButtonItem) SetCollection(value *TCollection) {
    TaskDialogRadioButtonItem_SetCollection(t.instance, CheckPtr(value))
}

func (t *TTaskDialogRadioButtonItem) Index() int32 {
    return TaskDialogRadioButtonItem_GetIndex(t.instance)
}

func (t *TTaskDialogRadioButtonItem) SetIndex(value int32) {
    TaskDialogRadioButtonItem_SetIndex(t.instance, value)
}

func (t *TTaskDialogRadioButtonItem) DisplayName() string {
    return TaskDialogRadioButtonItem_GetDisplayName(t.instance)
}

func (t *TTaskDialogRadioButtonItem) SetDisplayName(value string) {
    TaskDialogRadioButtonItem_SetDisplayName(t.instance, value)
}

