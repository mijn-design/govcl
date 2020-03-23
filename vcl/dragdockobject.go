
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

type TDragDockObject struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDragDockObject() *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = DragDockObject_Create()
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsDragDockObject(obj interface{}) *TDragDockObject {
    d := new(TDragDockObject)
    d.instance, d.ptr = getInstance(obj)
    if d.instance == 0 { return nil }
    return d
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsDragDockObject.
func DragDockObjectFromInst(inst uintptr) *TDragDockObject {
    return AsDragDockObject(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsDragDockObject.
func DragDockObjectFromObj(obj IObject) *TDragDockObject {
    return AsDragDockObject(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsDragDockObject.
func DragDockObjectFromUnsafePointer(ptr unsafe.Pointer) *TDragDockObject {
    return AsDragDockObject(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (d *TDragDockObject) Free() {
    if d.instance != 0 {
        DragDockObject_Free(d.instance)
        d.instance, d.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDragDockObject) Instance() uintptr {
    return d.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDragDockObject) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDragDockObject) IsValid() bool {
    return d.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (d *TDragDockObject) Is() TIs {
    return TIs(d.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (d *TDragDockObject) As() TAs {
//    return TAs(d.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDragDockObjectClass() TClass {
    return DragDockObject_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (d *TDragDockObject) Assign(Source *TDragObject) {
    DragDockObject_Assign(d.instance, CheckPtr(Source))
}

func (d *TDragDockObject) HideDragImage() {
    DragDockObject_HideDragImage(d.instance)
}

func (d *TDragDockObject) ShowDragImage() {
    DragDockObject_ShowDragImage(d.instance)
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDragDockObject) DisposeOf() {
    DragDockObject_DisposeOf(d.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDragDockObject) ClassType() TClass {
    return DragDockObject_ClassType(d.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDragDockObject) ClassName() string {
    return DragDockObject_ClassName(d.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDragDockObject) InstanceSize() int32 {
    return DragDockObject_InstanceSize(d.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDragDockObject) InheritsFrom(AClass TClass) bool {
    return DragDockObject_InheritsFrom(d.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDragDockObject) Equals(Obj IObject) bool {
    return DragDockObject_Equals(d.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDragDockObject) GetHashCode() int32 {
    return DragDockObject_GetHashCode(d.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (d *TDragDockObject) ToString() string {
    return DragDockObject_ToString(d.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (d *TDragDockObject) Brush() *TBrush {
    return AsBrush(DragDockObject_GetBrush(d.instance))
}

// CN: 设置画刷对象。
// EN: Set Brush.
func (d *TDragDockObject) SetBrush(value *TBrush) {
    DragDockObject_SetBrush(d.instance, CheckPtr(value))
}

func (d *TDragDockObject) DockRect() TRect {
    return DragDockObject_GetDockRect(d.instance)
}

func (d *TDragDockObject) SetDockRect(value TRect) {
    DragDockObject_SetDockRect(d.instance, value)
}

func (d *TDragDockObject) DropAlign() TAlign {
    return DragDockObject_GetDropAlign(d.instance)
}

func (d *TDragDockObject) DropOnControl() *TControl {
    return AsControl(DragDockObject_GetDropOnControl(d.instance))
}

func (d *TDragDockObject) EraseDockRect() TRect {
    return DragDockObject_GetEraseDockRect(d.instance)
}

func (d *TDragDockObject) SetEraseDockRect(value TRect) {
    DragDockObject_SetEraseDockRect(d.instance, value)
}

func (d *TDragDockObject) EraseWhenMoving() bool {
    return DragDockObject_GetEraseWhenMoving(d.instance)
}

func (d *TDragDockObject) Floating() bool {
    return DragDockObject_GetFloating(d.instance)
}

func (d *TDragDockObject) SetFloating(value bool) {
    DragDockObject_SetFloating(d.instance, value)
}

func (d *TDragDockObject) FrameWidth() int32 {
    return DragDockObject_GetFrameWidth(d.instance)
}

func (d *TDragDockObject) Control() *TControl {
    return AsControl(DragDockObject_GetControl(d.instance))
}

func (d *TDragDockObject) SetControl(value IControl) {
    DragDockObject_SetControl(d.instance, CheckPtr(value))
}

func (d *TDragDockObject) AlwaysShowDragImages() bool {
    return DragDockObject_GetAlwaysShowDragImages(d.instance)
}

func (d *TDragDockObject) SetAlwaysShowDragImages(value bool) {
    DragDockObject_SetAlwaysShowDragImages(d.instance, value)
}

func (d *TDragDockObject) Cancelling() bool {
    return DragDockObject_GetCancelling(d.instance)
}

func (d *TDragDockObject) SetCancelling(value bool) {
    DragDockObject_SetCancelling(d.instance, value)
}

func (d *TDragDockObject) DragHandle() HWND {
    return DragDockObject_GetDragHandle(d.instance)
}

func (d *TDragDockObject) SetDragHandle(value HWND) {
    DragDockObject_SetDragHandle(d.instance, value)
}

func (d *TDragDockObject) DragPos() TPoint {
    return DragDockObject_GetDragPos(d.instance)
}

func (d *TDragDockObject) SetDragPos(value TPoint) {
    DragDockObject_SetDragPos(d.instance, value)
}

func (d *TDragDockObject) DragTarget() uintptr {
    return DragDockObject_GetDragTarget(d.instance)
}

func (d *TDragDockObject) SetDragTarget(value uintptr) {
    DragDockObject_SetDragTarget(d.instance, value)
}

func (d *TDragDockObject) DragTargetPos() TPoint {
    return DragDockObject_GetDragTargetPos(d.instance)
}

func (d *TDragDockObject) SetDragTargetPos(value TPoint) {
    DragDockObject_SetDragTargetPos(d.instance, value)
}

func (d *TDragDockObject) Dropped() bool {
    return DragDockObject_GetDropped(d.instance)
}

func (d *TDragDockObject) MouseDeltaX() float64 {
    return DragDockObject_GetMouseDeltaX(d.instance)
}

func (d *TDragDockObject) MouseDeltaY() float64 {
    return DragDockObject_GetMouseDeltaY(d.instance)
}

func (d *TDragDockObject) RightClickCancels() bool {
    return DragDockObject_GetRightClickCancels(d.instance)
}

func (d *TDragDockObject) SetRightClickCancels(value bool) {
    DragDockObject_SetRightClickCancels(d.instance, value)
}

