
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

type TFlowPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewFlowPanel(owner IComponent) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = FlowPanel_Create(CheckPtr(owner))
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsFlowPanel(obj interface{}) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance, f.ptr = getInstance(obj)
    if f.instance == 0 { return nil }
    return f
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsFlowPanel.
func FlowPanelFromInst(inst uintptr) *TFlowPanel {
    return AsFlowPanel(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsFlowPanel.
func FlowPanelFromObj(obj IObject) *TFlowPanel {
    return AsFlowPanel(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsFlowPanel.
func FlowPanelFromUnsafePointer(ptr unsafe.Pointer) *TFlowPanel {
    return AsFlowPanel(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (f *TFlowPanel) Free() {
    if f.instance != 0 {
        FlowPanel_Free(f.instance)
        f.instance, f.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (f *TFlowPanel) Instance() uintptr {
    return f.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (f *TFlowPanel) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (f *TFlowPanel) IsValid() bool {
    return f.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (f *TFlowPanel) Is() TIs {
    return TIs(f.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (f *TFlowPanel) As() TAs {
//    return TAs(f.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TFlowPanelClass() TClass {
    return FlowPanel_StaticClassType()
}

func (f *TFlowPanel) GetControlIndex(AControl IControl) int32 {
    return FlowPanel_GetControlIndex(f.instance, CheckPtr(AControl))
}

func (f *TFlowPanel) SetControlIndex(AControl IControl, Index int32) {
    FlowPanel_SetControlIndex(f.instance, CheckPtr(AControl), Index)
}

// CN: 是否可以获得焦点。
// EN: .
func (f *TFlowPanel) CanFocus() bool {
    return FlowPanel_CanFocus(f.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (f *TFlowPanel) ContainsControl(Control IControl) bool {
    return FlowPanel_ContainsControl(f.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (f *TFlowPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(FlowPanel_ControlAtPos(f.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (f *TFlowPanel) DisableAlign() {
    FlowPanel_DisableAlign(f.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (f *TFlowPanel) EnableAlign() {
    FlowPanel_EnableAlign(f.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (f *TFlowPanel) FindChildControl(ControlName string) *TControl {
    return AsControl(FlowPanel_FindChildControl(f.instance, ControlName))
}

func (f *TFlowPanel) FlipChildren(AllLevels bool) {
    FlowPanel_FlipChildren(f.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (f *TFlowPanel) Focused() bool {
    return FlowPanel_Focused(f.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (f *TFlowPanel) HandleAllocated() bool {
    return FlowPanel_HandleAllocated(f.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (f *TFlowPanel) InsertControl(AControl IControl) {
    FlowPanel_InsertControl(f.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (f *TFlowPanel) Invalidate() {
    FlowPanel_Invalidate(f.instance)
}

// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (f *TFlowPanel) PaintTo(DC HDC, X int32, Y int32) {
    FlowPanel_PaintTo(f.instance, DC , X , Y)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (f *TFlowPanel) RemoveControl(AControl IControl) {
    FlowPanel_RemoveControl(f.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (f *TFlowPanel) Realign() {
    FlowPanel_Realign(f.instance)
}

// CN: 重绘。
// EN: Repaint.
func (f *TFlowPanel) Repaint() {
    FlowPanel_Repaint(f.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (f *TFlowPanel) ScaleBy(M int32, D int32) {
    FlowPanel_ScaleBy(f.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (f *TFlowPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    FlowPanel_ScrollBy(f.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (f *TFlowPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    FlowPanel_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (f *TFlowPanel) SetFocus() {
    FlowPanel_SetFocus(f.instance)
}

// CN: 控件更新。
// EN: Update.
func (f *TFlowPanel) Update() {
    FlowPanel_Update(f.instance)
}

// CN: 更新控件状态。
// EN: Update control status.
func (f *TFlowPanel) UpdateControlState() {
    FlowPanel_UpdateControlState(f.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (f *TFlowPanel) BringToFront() {
    FlowPanel_BringToFront(f.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (f *TFlowPanel) ClientToScreen(Point TPoint) TPoint {
    return FlowPanel_ClientToScreen(f.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (f *TFlowPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ClientToParent(f.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (f *TFlowPanel) Dragging() bool {
    return FlowPanel_Dragging(f.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (f *TFlowPanel) HasParent() bool {
    return FlowPanel_HasParent(f.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (f *TFlowPanel) Hide() {
    FlowPanel_Hide(f.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (f *TFlowPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return FlowPanel_Perform(f.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (f *TFlowPanel) Refresh() {
    FlowPanel_Refresh(f.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (f *TFlowPanel) ScreenToClient(Point TPoint) TPoint {
    return FlowPanel_ScreenToClient(f.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (f *TFlowPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ParentToClient(f.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (f *TFlowPanel) SendToBack() {
    FlowPanel_SendToBack(f.instance)
}

// CN: 显示控件。
// EN: Show control.
func (f *TFlowPanel) Show() {
    FlowPanel_Show(f.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (f *TFlowPanel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return FlowPanel_GetTextBuf(f.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (f *TFlowPanel) GetTextLen() int32 {
    return FlowPanel_GetTextLen(f.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (f *TFlowPanel) SetTextBuf(Buffer string) {
    FlowPanel_SetTextBuf(f.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (f *TFlowPanel) FindComponent(AName string) *TComponent {
    return AsComponent(FlowPanel_FindComponent(f.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (f *TFlowPanel) GetNamePath() string {
    return FlowPanel_GetNamePath(f.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (f *TFlowPanel) Assign(Source IObject) {
    FlowPanel_Assign(f.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (f *TFlowPanel) DisposeOf() {
    FlowPanel_DisposeOf(f.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (f *TFlowPanel) ClassType() TClass {
    return FlowPanel_ClassType(f.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (f *TFlowPanel) ClassName() string {
    return FlowPanel_ClassName(f.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (f *TFlowPanel) InstanceSize() int32 {
    return FlowPanel_InstanceSize(f.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (f *TFlowPanel) InheritsFrom(AClass TClass) bool {
    return FlowPanel_InheritsFrom(f.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (f *TFlowPanel) Equals(Obj IObject) bool {
    return FlowPanel_Equals(f.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (f *TFlowPanel) GetHashCode() int32 {
    return FlowPanel_GetHashCode(f.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (f *TFlowPanel) ToString() string {
    return FlowPanel_ToString(f.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (f *TFlowPanel) Align() TAlign {
    return FlowPanel_GetAlign(f.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (f *TFlowPanel) SetAlign(value TAlign) {
    FlowPanel_SetAlign(f.instance, value)
}

// CN: 获取文字对齐。
// EN: Get Text alignment.
func (f *TFlowPanel) Alignment() TAlignment {
    return FlowPanel_GetAlignment(f.instance)
}

// CN: 设置文字对齐。
// EN: Set Text alignment.
func (f *TFlowPanel) SetAlignment(value TAlignment) {
    FlowPanel_SetAlignment(f.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (f *TFlowPanel) Anchors() TAnchors {
    return FlowPanel_GetAnchors(f.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (f *TFlowPanel) SetAnchors(value TAnchors) {
    FlowPanel_SetAnchors(f.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (f *TFlowPanel) AutoSize() bool {
    return FlowPanel_GetAutoSize(f.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (f *TFlowPanel) SetAutoSize(value bool) {
    FlowPanel_SetAutoSize(f.instance, value)
}

func (f *TFlowPanel) AutoWrap() bool {
    return FlowPanel_GetAutoWrap(f.instance)
}

func (f *TFlowPanel) SetAutoWrap(value bool) {
    FlowPanel_SetAutoWrap(f.instance, value)
}

func (f *TFlowPanel) BevelEdges() TBevelEdges {
    return FlowPanel_GetBevelEdges(f.instance)
}

func (f *TFlowPanel) SetBevelEdges(value TBevelEdges) {
    FlowPanel_SetBevelEdges(f.instance, value)
}

func (f *TFlowPanel) BevelInner() TBevelCut {
    return FlowPanel_GetBevelInner(f.instance)
}

func (f *TFlowPanel) SetBevelInner(value TBevelCut) {
    FlowPanel_SetBevelInner(f.instance, value)
}

func (f *TFlowPanel) BevelKind() TBevelKind {
    return FlowPanel_GetBevelKind(f.instance)
}

func (f *TFlowPanel) SetBevelKind(value TBevelKind) {
    FlowPanel_SetBevelKind(f.instance, value)
}

func (f *TFlowPanel) BevelOuter() TBevelCut {
    return FlowPanel_GetBevelOuter(f.instance)
}

func (f *TFlowPanel) SetBevelOuter(value TBevelCut) {
    FlowPanel_SetBevelOuter(f.instance, value)
}

func (f *TFlowPanel) BiDiMode() TBiDiMode {
    return FlowPanel_GetBiDiMode(f.instance)
}

func (f *TFlowPanel) SetBiDiMode(value TBiDiMode) {
    FlowPanel_SetBiDiMode(f.instance, value)
}

// CN: 获取边框的宽度。
// EN: .
func (f *TFlowPanel) BorderWidth() int32 {
    return FlowPanel_GetBorderWidth(f.instance)
}

// CN: 设置边框的宽度。
// EN: .
func (f *TFlowPanel) SetBorderWidth(value int32) {
    FlowPanel_SetBorderWidth(f.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (f *TFlowPanel) BorderStyle() TBorderStyle {
    return FlowPanel_GetBorderStyle(f.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (f *TFlowPanel) SetBorderStyle(value TBorderStyle) {
    FlowPanel_SetBorderStyle(f.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (f *TFlowPanel) Caption() string {
    return FlowPanel_GetCaption(f.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (f *TFlowPanel) SetCaption(value string) {
    FlowPanel_SetCaption(f.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (f *TFlowPanel) Color() TColor {
    return FlowPanel_GetColor(f.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (f *TFlowPanel) SetColor(value TColor) {
    FlowPanel_SetColor(f.instance, value)
}

func (f *TFlowPanel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(FlowPanel_GetConstraints(f.instance))
}

func (f *TFlowPanel) SetConstraints(value *TSizeConstraints) {
    FlowPanel_SetConstraints(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) Ctl3D() bool {
    return FlowPanel_GetCtl3D(f.instance)
}

func (f *TFlowPanel) SetCtl3D(value bool) {
    FlowPanel_SetCtl3D(f.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (f *TFlowPanel) UseDockManager() bool {
    return FlowPanel_GetUseDockManager(f.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (f *TFlowPanel) SetUseDockManager(value bool) {
    FlowPanel_SetUseDockManager(f.instance, value)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (f *TFlowPanel) DockSite() bool {
    return FlowPanel_GetDockSite(f.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (f *TFlowPanel) SetDockSite(value bool) {
    FlowPanel_SetDockSite(f.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (f *TFlowPanel) DoubleBuffered() bool {
    return FlowPanel_GetDoubleBuffered(f.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (f *TFlowPanel) SetDoubleBuffered(value bool) {
    FlowPanel_SetDoubleBuffered(f.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (f *TFlowPanel) DragCursor() TCursor {
    return FlowPanel_GetDragCursor(f.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (f *TFlowPanel) SetDragCursor(value TCursor) {
    FlowPanel_SetDragCursor(f.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (f *TFlowPanel) DragKind() TDragKind {
    return FlowPanel_GetDragKind(f.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (f *TFlowPanel) SetDragKind(value TDragKind) {
    FlowPanel_SetDragKind(f.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (f *TFlowPanel) DragMode() TDragMode {
    return FlowPanel_GetDragMode(f.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (f *TFlowPanel) SetDragMode(value TDragMode) {
    FlowPanel_SetDragMode(f.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (f *TFlowPanel) Enabled() bool {
    return FlowPanel_GetEnabled(f.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (f *TFlowPanel) SetEnabled(value bool) {
    FlowPanel_SetEnabled(f.instance, value)
}

func (f *TFlowPanel) FlowStyle() TFlowStyle {
    return FlowPanel_GetFlowStyle(f.instance)
}

func (f *TFlowPanel) SetFlowStyle(value TFlowStyle) {
    FlowPanel_SetFlowStyle(f.instance, value)
}

func (f *TFlowPanel) FullRepaint() bool {
    return FlowPanel_GetFullRepaint(f.instance)
}

func (f *TFlowPanel) SetFullRepaint(value bool) {
    FlowPanel_SetFullRepaint(f.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (f *TFlowPanel) Font() *TFont {
    return AsFont(FlowPanel_GetFont(f.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (f *TFlowPanel) SetFont(value *TFont) {
    FlowPanel_SetFont(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) Locked() bool {
    return FlowPanel_GetLocked(f.instance)
}

func (f *TFlowPanel) SetLocked(value bool) {
    FlowPanel_SetLocked(f.instance, value)
}

func (f *TFlowPanel) ParentBackground() bool {
    return FlowPanel_GetParentBackground(f.instance)
}

func (f *TFlowPanel) SetParentBackground(value bool) {
    FlowPanel_SetParentBackground(f.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (f *TFlowPanel) ParentColor() bool {
    return FlowPanel_GetParentColor(f.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (f *TFlowPanel) SetParentColor(value bool) {
    FlowPanel_SetParentColor(f.instance, value)
}

func (f *TFlowPanel) ParentCtl3D() bool {
    return FlowPanel_GetParentCtl3D(f.instance)
}

func (f *TFlowPanel) SetParentCtl3D(value bool) {
    FlowPanel_SetParentCtl3D(f.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (f *TFlowPanel) ParentDoubleBuffered() bool {
    return FlowPanel_GetParentDoubleBuffered(f.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (f *TFlowPanel) SetParentDoubleBuffered(value bool) {
    FlowPanel_SetParentDoubleBuffered(f.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (f *TFlowPanel) ParentFont() bool {
    return FlowPanel_GetParentFont(f.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (f *TFlowPanel) SetParentFont(value bool) {
    FlowPanel_SetParentFont(f.instance, value)
}

func (f *TFlowPanel) ParentShowHint() bool {
    return FlowPanel_GetParentShowHint(f.instance)
}

func (f *TFlowPanel) SetParentShowHint(value bool) {
    FlowPanel_SetParentShowHint(f.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (f *TFlowPanel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(FlowPanel_GetPopupMenu(f.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (f *TFlowPanel) SetPopupMenu(value IComponent) {
    FlowPanel_SetPopupMenu(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) ShowCaption() bool {
    return FlowPanel_GetShowCaption(f.instance)
}

func (f *TFlowPanel) SetShowCaption(value bool) {
    FlowPanel_SetShowCaption(f.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (f *TFlowPanel) ShowHint() bool {
    return FlowPanel_GetShowHint(f.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (f *TFlowPanel) SetShowHint(value bool) {
    FlowPanel_SetShowHint(f.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (f *TFlowPanel) TabOrder() TTabOrder {
    return FlowPanel_GetTabOrder(f.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (f *TFlowPanel) SetTabOrder(value TTabOrder) {
    FlowPanel_SetTabOrder(f.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (f *TFlowPanel) TabStop() bool {
    return FlowPanel_GetTabStop(f.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (f *TFlowPanel) SetTabStop(value bool) {
    FlowPanel_SetTabStop(f.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (f *TFlowPanel) Visible() bool {
    return FlowPanel_GetVisible(f.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (f *TFlowPanel) SetVisible(value bool) {
    FlowPanel_SetVisible(f.instance, value)
}

// CN: 获取样式元素。
// EN: Get Style element.
func (f *TFlowPanel) StyleElements() TStyleElements {
    return FlowPanel_GetStyleElements(f.instance)
}

// CN: 设置样式元素。
// EN: Set Style element.
func (f *TFlowPanel) SetStyleElements(value TStyleElements) {
    FlowPanel_SetStyleElements(f.instance, value)
}

func (f *TFlowPanel) SetOnAlignPosition(fn TAlignPositionEvent) {
    FlowPanel_SetOnAlignPosition(f.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (f *TFlowPanel) SetOnClick(fn TNotifyEvent) {
    FlowPanel_SetOnClick(f.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (f *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
    FlowPanel_SetOnContextPopup(f.instance, fn)
}

func (f *TFlowPanel) SetOnDockDrop(fn TDockDropEvent) {
    FlowPanel_SetOnDockDrop(f.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (f *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
    FlowPanel_SetOnDblClick(f.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (f *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
    FlowPanel_SetOnDragDrop(f.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (f *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
    FlowPanel_SetOnDragOver(f.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (f *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
    FlowPanel_SetOnEndDock(f.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (f *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
    FlowPanel_SetOnEndDrag(f.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (f *TFlowPanel) SetOnEnter(fn TNotifyEvent) {
    FlowPanel_SetOnEnter(f.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (f *TFlowPanel) SetOnExit(fn TNotifyEvent) {
    FlowPanel_SetOnExit(f.instance, fn)
}

func (f *TFlowPanel) SetOnGesture(fn TGestureEvent) {
    FlowPanel_SetOnGesture(f.instance, fn)
}

func (f *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    FlowPanel_SetOnGetSiteInfo(f.instance, fn)
}

func (f *TFlowPanel) SetOnMouseActivate(fn TMouseActivateEvent) {
    FlowPanel_SetOnMouseActivate(f.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (f *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
    FlowPanel_SetOnMouseDown(f.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (f *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
    FlowPanel_SetOnMouseEnter(f.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (f *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
    FlowPanel_SetOnMouseLeave(f.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (f *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    FlowPanel_SetOnMouseMove(f.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (f *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
    FlowPanel_SetOnMouseUp(f.instance, fn)
}

// CN: 设置大小被改变事件。
// EN: .
func (f *TFlowPanel) SetOnResize(fn TNotifyEvent) {
    FlowPanel_SetOnResize(f.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (f *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
    FlowPanel_SetOnStartDock(f.instance, fn)
}

func (f *TFlowPanel) SetOnUnDock(fn TUnDockEvent) {
    FlowPanel_SetOnUnDock(f.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (f *TFlowPanel) DockClientCount() int32 {
    return FlowPanel_GetDockClientCount(f.instance)
}

// CN: 获取禁用对齐。
// EN: .
func (f *TFlowPanel) AlignDisabled() bool {
    return FlowPanel_GetAlignDisabled(f.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (f *TFlowPanel) MouseInClient() bool {
    return FlowPanel_GetMouseInClient(f.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (f *TFlowPanel) VisibleDockClientCount() int32 {
    return FlowPanel_GetVisibleDockClientCount(f.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (f *TFlowPanel) Brush() *TBrush {
    return AsBrush(FlowPanel_GetBrush(f.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (f *TFlowPanel) ControlCount() int32 {
    return FlowPanel_GetControlCount(f.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (f *TFlowPanel) Handle() HWND {
    return FlowPanel_GetHandle(f.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (f *TFlowPanel) ParentWindow() HWND {
    return FlowPanel_GetParentWindow(f.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (f *TFlowPanel) SetParentWindow(value HWND) {
    FlowPanel_SetParentWindow(f.instance, value)
}

func (f *TFlowPanel) Action() *TAction {
    return AsAction(FlowPanel_GetAction(f.instance))
}

func (f *TFlowPanel) SetAction(value IComponent) {
    FlowPanel_SetAction(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) BoundsRect() TRect {
    return FlowPanel_GetBoundsRect(f.instance)
}

func (f *TFlowPanel) SetBoundsRect(value TRect) {
    FlowPanel_SetBoundsRect(f.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (f *TFlowPanel) ClientHeight() int32 {
    return FlowPanel_GetClientHeight(f.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (f *TFlowPanel) SetClientHeight(value int32) {
    FlowPanel_SetClientHeight(f.instance, value)
}

func (f *TFlowPanel) ClientOrigin() TPoint {
    return FlowPanel_GetClientOrigin(f.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (f *TFlowPanel) ClientRect() TRect {
    return FlowPanel_GetClientRect(f.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (f *TFlowPanel) ClientWidth() int32 {
    return FlowPanel_GetClientWidth(f.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (f *TFlowPanel) SetClientWidth(value int32) {
    FlowPanel_SetClientWidth(f.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (f *TFlowPanel) ControlState() TControlState {
    return FlowPanel_GetControlState(f.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (f *TFlowPanel) SetControlState(value TControlState) {
    FlowPanel_SetControlState(f.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (f *TFlowPanel) ControlStyle() TControlStyle {
    return FlowPanel_GetControlStyle(f.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (f *TFlowPanel) SetControlStyle(value TControlStyle) {
    FlowPanel_SetControlStyle(f.instance, value)
}

func (f *TFlowPanel) ExplicitLeft() int32 {
    return FlowPanel_GetExplicitLeft(f.instance)
}

func (f *TFlowPanel) ExplicitTop() int32 {
    return FlowPanel_GetExplicitTop(f.instance)
}

func (f *TFlowPanel) ExplicitWidth() int32 {
    return FlowPanel_GetExplicitWidth(f.instance)
}

func (f *TFlowPanel) ExplicitHeight() int32 {
    return FlowPanel_GetExplicitHeight(f.instance)
}

func (f *TFlowPanel) Floating() bool {
    return FlowPanel_GetFloating(f.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (f *TFlowPanel) Parent() *TWinControl {
    return AsWinControl(FlowPanel_GetParent(f.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (f *TFlowPanel) SetParent(value IWinControl) {
    FlowPanel_SetParent(f.instance, CheckPtr(value))
}

// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (f *TFlowPanel) AlignWithMargins() bool {
    return FlowPanel_GetAlignWithMargins(f.instance)
}

// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (f *TFlowPanel) SetAlignWithMargins(value bool) {
    FlowPanel_SetAlignWithMargins(f.instance, value)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (f *TFlowPanel) Left() int32 {
    return FlowPanel_GetLeft(f.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (f *TFlowPanel) SetLeft(value int32) {
    FlowPanel_SetLeft(f.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (f *TFlowPanel) Top() int32 {
    return FlowPanel_GetTop(f.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (f *TFlowPanel) SetTop(value int32) {
    FlowPanel_SetTop(f.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (f *TFlowPanel) Width() int32 {
    return FlowPanel_GetWidth(f.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (f *TFlowPanel) SetWidth(value int32) {
    FlowPanel_SetWidth(f.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (f *TFlowPanel) Height() int32 {
    return FlowPanel_GetHeight(f.instance)
}

// CN: 设置高度。
// EN: Set height.
func (f *TFlowPanel) SetHeight(value int32) {
    FlowPanel_SetHeight(f.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (f *TFlowPanel) Cursor() TCursor {
    return FlowPanel_GetCursor(f.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (f *TFlowPanel) SetCursor(value TCursor) {
    FlowPanel_SetCursor(f.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (f *TFlowPanel) Hint() string {
    return FlowPanel_GetHint(f.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (f *TFlowPanel) SetHint(value string) {
    FlowPanel_SetHint(f.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (f *TFlowPanel) Margins() *TMargins {
    return AsMargins(FlowPanel_GetMargins(f.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (f *TFlowPanel) SetMargins(value *TMargins) {
    FlowPanel_SetMargins(f.instance, CheckPtr(value))
}

// CN: 获取自定义提示。
// EN: Get custom hint.
func (f *TFlowPanel) CustomHint() *TCustomHint {
    return AsCustomHint(FlowPanel_GetCustomHint(f.instance))
}

// CN: 设置自定义提示。
// EN: Set custom hint.
func (f *TFlowPanel) SetCustomHint(value IComponent) {
    FlowPanel_SetCustomHint(f.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (f *TFlowPanel) ComponentCount() int32 {
    return FlowPanel_GetComponentCount(f.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (f *TFlowPanel) ComponentIndex() int32 {
    return FlowPanel_GetComponentIndex(f.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (f *TFlowPanel) SetComponentIndex(value int32) {
    FlowPanel_SetComponentIndex(f.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (f *TFlowPanel) Owner() *TComponent {
    return AsComponent(FlowPanel_GetOwner(f.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (f *TFlowPanel) Name() string {
    return FlowPanel_GetName(f.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (f *TFlowPanel) SetName(value string) {
    FlowPanel_SetName(f.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (f *TFlowPanel) Tag() int {
    return FlowPanel_GetTag(f.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (f *TFlowPanel) SetTag(value int) {
    FlowPanel_SetTag(f.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (f *TFlowPanel) DockClients(Index int32) *TControl {
    return AsControl(FlowPanel_GetDockClients(f.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (f *TFlowPanel) Controls(Index int32) *TControl {
    return AsControl(FlowPanel_GetControls(f.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (f *TFlowPanel) Components(AIndex int32) *TComponent {
    return AsComponent(FlowPanel_GetComponents(f.instance, AIndex))
}

