
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

type TTaskDialogProgressBar struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskDialogProgressBar() *TTaskDialogProgressBar {
    t := new(TTaskDialogProgressBar)
    t.instance = TaskDialogProgressBar_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTaskDialogProgressBar(obj interface{}) *TTaskDialogProgressBar {
    t := new(TTaskDialogProgressBar)
    t.instance, t.ptr = getInstance(obj)
    if t.instance == 0 { return nil }
    return t
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskDialogProgressBar.
func TaskDialogProgressBarFromInst(inst uintptr) *TTaskDialogProgressBar {
    return AsTaskDialogProgressBar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTaskDialogProgressBar.
func TaskDialogProgressBarFromObj(obj IObject) *TTaskDialogProgressBar {
    return AsTaskDialogProgressBar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskDialogProgressBar.
func TaskDialogProgressBarFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogProgressBar {
    return AsTaskDialogProgressBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TTaskDialogProgressBar) Free() {
    if t.instance != 0 {
        TaskDialogProgressBar_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialogProgressBar) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialogProgressBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialogProgressBar) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTaskDialogProgressBar) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTaskDialogProgressBar) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogProgressBarClass() TClass {
    return TaskDialogProgressBar_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialogProgressBar) Assign(Source IObject) {
    TaskDialogProgressBar_Assign(t.instance, CheckPtr(Source))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialogProgressBar) GetNamePath() string {
    return TaskDialogProgressBar_GetNamePath(t.instance)
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskDialogProgressBar) DisposeOf() {
    TaskDialogProgressBar_DisposeOf(t.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialogProgressBar) ClassType() TClass {
    return TaskDialogProgressBar_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialogProgressBar) ClassName() string {
    return TaskDialogProgressBar_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialogProgressBar) InstanceSize() int32 {
    return TaskDialogProgressBar_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialogProgressBar) InheritsFrom(AClass TClass) bool {
    return TaskDialogProgressBar_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialogProgressBar) Equals(Obj IObject) bool {
    return TaskDialogProgressBar_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialogProgressBar) GetHashCode() int32 {
    return TaskDialogProgressBar_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialogProgressBar) ToString() string {
    return TaskDialogProgressBar_ToString(t.instance)
}

func (t *TTaskDialogProgressBar) Max() int32 {
    return TaskDialogProgressBar_GetMax(t.instance)
}

func (t *TTaskDialogProgressBar) SetMax(value int32) {
    TaskDialogProgressBar_SetMax(t.instance, value)
}

func (t *TTaskDialogProgressBar) Min() int32 {
    return TaskDialogProgressBar_GetMin(t.instance)
}

func (t *TTaskDialogProgressBar) SetMin(value int32) {
    TaskDialogProgressBar_SetMin(t.instance, value)
}

func (t *TTaskDialogProgressBar) Position() int32 {
    return TaskDialogProgressBar_GetPosition(t.instance)
}

func (t *TTaskDialogProgressBar) SetPosition(value int32) {
    TaskDialogProgressBar_SetPosition(t.instance, value)
}

func (t *TTaskDialogProgressBar) State() TProgressBarState {
    return TaskDialogProgressBar_GetState(t.instance)
}

func (t *TTaskDialogProgressBar) SetState(value TProgressBarState) {
    TaskDialogProgressBar_SetState(t.instance, value)
}

