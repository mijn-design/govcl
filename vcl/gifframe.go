
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

type TGIFFrame struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGIFFrame() *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = GIFFrame_Create()
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsGIFFrame(obj interface{}) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance, g.ptr = getInstance(obj)
    if g.instance == 0 { return nil }
    return g
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsGIFFrame.
func GIFFrameFromInst(inst uintptr) *TGIFFrame {
    return AsGIFFrame(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsGIFFrame.
func GIFFrameFromObj(obj IObject) *TGIFFrame {
    return AsGIFFrame(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsGIFFrame.
func GIFFrameFromUnsafePointer(ptr unsafe.Pointer) *TGIFFrame {
    return AsGIFFrame(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (g *TGIFFrame) Free() {
    if g.instance != 0 {
        GIFFrame_Free(g.instance)
        g.instance, g.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGIFFrame) Instance() uintptr {
    return g.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGIFFrame) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGIFFrame) IsValid() bool {
    return g.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (g *TGIFFrame) Is() TIs {
    return TIs(g.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (g *TGIFFrame) As() TAs {
//    return TAs(g.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGIFFrameClass() TClass {
    return GIFFrame_StaticClassType()
}

// CN: 清除。
// EN: .
func (g *TGIFFrame) Clear() {
    GIFFrame_Clear(g.instance)
}

// CN: 保存至流。
// EN: .
func (g *TGIFFrame) SaveToStream(Stream IObject) {
    GIFFrame_SaveToStream(g.instance, CheckPtr(Stream))
}

// CN: 文件流加载。
// EN: .
func (g *TGIFFrame) LoadFromStream(Stream IObject) {
    GIFFrame_LoadFromStream(g.instance, CheckPtr(Stream))
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGIFFrame) Assign(Source IObject) {
    GIFFrame_Assign(g.instance, CheckPtr(Source))
}

// CN: 保存至文件。
// EN: .
func (g *TGIFFrame) SaveToFile(Filename string) {
    GIFFrame_SaveToFile(g.instance, Filename)
}

// CN: 从文件加载。
// EN: .
func (g *TGIFFrame) LoadFromFile(Filename string) {
    GIFFrame_LoadFromFile(g.instance, Filename)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGIFFrame) GetNamePath() string {
    return GIFFrame_GetNamePath(g.instance)
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGIFFrame) DisposeOf() {
    GIFFrame_DisposeOf(g.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGIFFrame) ClassType() TClass {
    return GIFFrame_ClassType(g.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGIFFrame) ClassName() string {
    return GIFFrame_ClassName(g.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGIFFrame) InstanceSize() int32 {
    return GIFFrame_InstanceSize(g.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGIFFrame) InheritsFrom(AClass TClass) bool {
    return GIFFrame_InheritsFrom(g.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGIFFrame) Equals(Obj IObject) bool {
    return GIFFrame_Equals(g.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGIFFrame) GetHashCode() int32 {
    return GIFFrame_GetHashCode(g.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (g *TGIFFrame) ToString() string {
    return GIFFrame_ToString(g.instance)
}

func (g *TGIFFrame) HasBitmap() bool {
    return GIFFrame_GetHasBitmap(g.instance)
}

func (g *TGIFFrame) SetHasBitmap(value bool) {
    GIFFrame_SetHasBitmap(g.instance, value)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (g *TGIFFrame) Left() uint16 {
    return GIFFrame_GetLeft(g.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (g *TGIFFrame) SetLeft(value uint16) {
    GIFFrame_SetLeft(g.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (g *TGIFFrame) Top() uint16 {
    return GIFFrame_GetTop(g.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (g *TGIFFrame) SetTop(value uint16) {
    GIFFrame_SetTop(g.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (g *TGIFFrame) Width() uint16 {
    return GIFFrame_GetWidth(g.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (g *TGIFFrame) SetWidth(value uint16) {
    GIFFrame_SetWidth(g.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (g *TGIFFrame) Height() uint16 {
    return GIFFrame_GetHeight(g.instance)
}

// CN: 设置高度。
// EN: Set height.
func (g *TGIFFrame) SetHeight(value uint16) {
    GIFFrame_SetHeight(g.instance, value)
}

func (g *TGIFFrame) BoundsRect() TRect {
    return GIFFrame_GetBoundsRect(g.instance)
}

func (g *TGIFFrame) SetBoundsRect(value TRect) {
    GIFFrame_SetBoundsRect(g.instance, value)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (g *TGIFFrame) ClientRect() TRect {
    return GIFFrame_GetClientRect(g.instance)
}

func (g *TGIFFrame) Data() uintptr {
    return GIFFrame_GetData(g.instance)
}

func (g *TGIFFrame) DataSize() int32 {
    return GIFFrame_GetDataSize(g.instance)
}

func (g *TGIFFrame) Version() TGIFVersion {
    return GIFFrame_GetVersion(g.instance)
}

func (g *TGIFFrame) BitsPerPixel() int32 {
    return GIFFrame_GetBitsPerPixel(g.instance)
}

func (g *TGIFFrame) Bitmap() *TBitmap {
    return AsBitmap(GIFFrame_GetBitmap(g.instance))
}

func (g *TGIFFrame) SetBitmap(value *TBitmap) {
    GIFFrame_SetBitmap(g.instance, CheckPtr(value))
}

func (g *TGIFFrame) Palette() HPALETTE {
    return GIFFrame_GetPalette(g.instance)
}

func (g *TGIFFrame) SetPalette(value HPALETTE) {
    GIFFrame_SetPalette(g.instance, value)
}

func (g *TGIFFrame) Empty() bool {
    return GIFFrame_GetEmpty(g.instance)
}

// CN: 获取透明。
// EN: Get transparent.
func (g *TGIFFrame) Transparent() bool {
    return GIFFrame_GetTransparent(g.instance)
}

