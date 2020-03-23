
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

type TScrollBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewScrollBar(owner IComponent) *TScrollBar {
    s := new(TScrollBar)
    s.instance = ScrollBar_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsScrollBar(obj interface{}) *TScrollBar {
    s := new(TScrollBar)
    s.instance, s.ptr = getInstance(obj)
    if s.instance == 0 { return nil }
    return s
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsScrollBar.
func ScrollBarFromInst(inst uintptr) *TScrollBar {
    return AsScrollBar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsScrollBar.
func ScrollBarFromObj(obj IObject) *TScrollBar {
    return AsScrollBar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsScrollBar.
func ScrollBarFromUnsafePointer(ptr unsafe.Pointer) *TScrollBar {
    return AsScrollBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (s *TScrollBar) Free() {
    if s.instance != 0 {
        ScrollBar_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TScrollBar) Instance() uintptr {
    return s.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TScrollBar) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TScrollBar) IsValid() bool {
    return s.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (s *TScrollBar) Is() TIs {
    return TIs(s.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (s *TScrollBar) As() TAs {
//    return TAs(s.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TScrollBarClass() TClass {
    return ScrollBar_StaticClassType()
}

func (s *TScrollBar) SetParams(APosition int32, AMin int32, AMax int32) {
    ScrollBar_SetParams(s.instance, APosition , AMin , AMax)
}

// CN: 是否可以获得焦点。
// EN: .
func (s *TScrollBar) CanFocus() bool {
    return ScrollBar_CanFocus(s.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TScrollBar) ContainsControl(Control IControl) bool {
    return ScrollBar_ContainsControl(s.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TScrollBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ScrollBar_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TScrollBar) DisableAlign() {
    ScrollBar_DisableAlign(s.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TScrollBar) EnableAlign() {
    ScrollBar_EnableAlign(s.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (s *TScrollBar) FindChildControl(ControlName string) *TControl {
    return AsControl(ScrollBar_FindChildControl(s.instance, ControlName))
}

func (s *TScrollBar) FlipChildren(AllLevels bool) {
    ScrollBar_FlipChildren(s.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TScrollBar) Focused() bool {
    return ScrollBar_Focused(s.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TScrollBar) HandleAllocated() bool {
    return ScrollBar_HandleAllocated(s.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (s *TScrollBar) InsertControl(AControl IControl) {
    ScrollBar_InsertControl(s.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (s *TScrollBar) Invalidate() {
    ScrollBar_Invalidate(s.instance)
}

// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (s *TScrollBar) PaintTo(DC HDC, X int32, Y int32) {
    ScrollBar_PaintTo(s.instance, DC , X , Y)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (s *TScrollBar) RemoveControl(AControl IControl) {
    ScrollBar_RemoveControl(s.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (s *TScrollBar) Realign() {
    ScrollBar_Realign(s.instance)
}

// CN: 重绘。
// EN: Repaint.
func (s *TScrollBar) Repaint() {
    ScrollBar_Repaint(s.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (s *TScrollBar) ScaleBy(M int32, D int32) {
    ScrollBar_ScaleBy(s.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TScrollBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ScrollBar_ScrollBy(s.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TScrollBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ScrollBar_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TScrollBar) SetFocus() {
    ScrollBar_SetFocus(s.instance)
}

// CN: 控件更新。
// EN: Update.
func (s *TScrollBar) Update() {
    ScrollBar_Update(s.instance)
}

// CN: 更新控件状态。
// EN: Update control status.
func (s *TScrollBar) UpdateControlState() {
    ScrollBar_UpdateControlState(s.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TScrollBar) BringToFront() {
    ScrollBar_BringToFront(s.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TScrollBar) ClientToScreen(Point TPoint) TPoint {
    return ScrollBar_ClientToScreen(s.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TScrollBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBar_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TScrollBar) Dragging() bool {
    return ScrollBar_Dragging(s.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TScrollBar) HasParent() bool {
    return ScrollBar_HasParent(s.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (s *TScrollBar) Hide() {
    ScrollBar_Hide(s.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (s *TScrollBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ScrollBar_Perform(s.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (s *TScrollBar) Refresh() {
    ScrollBar_Refresh(s.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TScrollBar) ScreenToClient(Point TPoint) TPoint {
    return ScrollBar_ScreenToClient(s.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TScrollBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBar_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TScrollBar) SendToBack() {
    ScrollBar_SendToBack(s.instance)
}

// CN: 显示控件。
// EN: Show control.
func (s *TScrollBar) Show() {
    ScrollBar_Show(s.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TScrollBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ScrollBar_GetTextBuf(s.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TScrollBar) GetTextLen() int32 {
    return ScrollBar_GetTextLen(s.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TScrollBar) SetTextBuf(Buffer string) {
    ScrollBar_SetTextBuf(s.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TScrollBar) FindComponent(AName string) *TComponent {
    return AsComponent(ScrollBar_FindComponent(s.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TScrollBar) GetNamePath() string {
    return ScrollBar_GetNamePath(s.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TScrollBar) Assign(Source IObject) {
    ScrollBar_Assign(s.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TScrollBar) DisposeOf() {
    ScrollBar_DisposeOf(s.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TScrollBar) ClassType() TClass {
    return ScrollBar_ClassType(s.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TScrollBar) ClassName() string {
    return ScrollBar_ClassName(s.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TScrollBar) InstanceSize() int32 {
    return ScrollBar_InstanceSize(s.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TScrollBar) InheritsFrom(AClass TClass) bool {
    return ScrollBar_InheritsFrom(s.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TScrollBar) Equals(Obj IObject) bool {
    return ScrollBar_Equals(s.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TScrollBar) GetHashCode() int32 {
    return ScrollBar_GetHashCode(s.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (s *TScrollBar) ToString() string {
    return ScrollBar_ToString(s.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TScrollBar) Align() TAlign {
    return ScrollBar_GetAlign(s.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TScrollBar) SetAlign(value TAlign) {
    ScrollBar_SetAlign(s.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (s *TScrollBar) Anchors() TAnchors {
    return ScrollBar_GetAnchors(s.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (s *TScrollBar) SetAnchors(value TAnchors) {
    ScrollBar_SetAnchors(s.instance, value)
}

func (s *TScrollBar) BiDiMode() TBiDiMode {
    return ScrollBar_GetBiDiMode(s.instance)
}

func (s *TScrollBar) SetBiDiMode(value TBiDiMode) {
    ScrollBar_SetBiDiMode(s.instance, value)
}

func (s *TScrollBar) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ScrollBar_GetConstraints(s.instance))
}

func (s *TScrollBar) SetConstraints(value *TSizeConstraints) {
    ScrollBar_SetConstraints(s.instance, CheckPtr(value))
}

func (s *TScrollBar) Ctl3D() bool {
    return ScrollBar_GetCtl3D(s.instance)
}

func (s *TScrollBar) SetCtl3D(value bool) {
    ScrollBar_SetCtl3D(s.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TScrollBar) DoubleBuffered() bool {
    return ScrollBar_GetDoubleBuffered(s.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TScrollBar) SetDoubleBuffered(value bool) {
    ScrollBar_SetDoubleBuffered(s.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (s *TScrollBar) DragCursor() TCursor {
    return ScrollBar_GetDragCursor(s.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (s *TScrollBar) SetDragCursor(value TCursor) {
    ScrollBar_SetDragCursor(s.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (s *TScrollBar) DragKind() TDragKind {
    return ScrollBar_GetDragKind(s.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (s *TScrollBar) SetDragKind(value TDragKind) {
    ScrollBar_SetDragKind(s.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (s *TScrollBar) DragMode() TDragMode {
    return ScrollBar_GetDragMode(s.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (s *TScrollBar) SetDragMode(value TDragMode) {
    ScrollBar_SetDragMode(s.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TScrollBar) Enabled() bool {
    return ScrollBar_GetEnabled(s.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TScrollBar) SetEnabled(value bool) {
    ScrollBar_SetEnabled(s.instance, value)
}

func (s *TScrollBar) Kind() TScrollBarKind {
    return ScrollBar_GetKind(s.instance)
}

func (s *TScrollBar) SetKind(value TScrollBarKind) {
    ScrollBar_SetKind(s.instance, value)
}

func (s *TScrollBar) LargeChange() TScrollBarInc {
    return ScrollBar_GetLargeChange(s.instance)
}

func (s *TScrollBar) SetLargeChange(value TScrollBarInc) {
    ScrollBar_SetLargeChange(s.instance, value)
}

func (s *TScrollBar) Max() int32 {
    return ScrollBar_GetMax(s.instance)
}

func (s *TScrollBar) SetMax(value int32) {
    ScrollBar_SetMax(s.instance, value)
}

func (s *TScrollBar) Min() int32 {
    return ScrollBar_GetMin(s.instance)
}

func (s *TScrollBar) SetMin(value int32) {
    ScrollBar_SetMin(s.instance, value)
}

func (s *TScrollBar) PageSize() int32 {
    return ScrollBar_GetPageSize(s.instance)
}

func (s *TScrollBar) SetPageSize(value int32) {
    ScrollBar_SetPageSize(s.instance, value)
}

func (s *TScrollBar) ParentCtl3D() bool {
    return ScrollBar_GetParentCtl3D(s.instance)
}

func (s *TScrollBar) SetParentCtl3D(value bool) {
    ScrollBar_SetParentCtl3D(s.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TScrollBar) ParentDoubleBuffered() bool {
    return ScrollBar_GetParentDoubleBuffered(s.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TScrollBar) SetParentDoubleBuffered(value bool) {
    ScrollBar_SetParentDoubleBuffered(s.instance, value)
}

func (s *TScrollBar) ParentShowHint() bool {
    return ScrollBar_GetParentShowHint(s.instance)
}

func (s *TScrollBar) SetParentShowHint(value bool) {
    ScrollBar_SetParentShowHint(s.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TScrollBar) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ScrollBar_GetPopupMenu(s.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TScrollBar) SetPopupMenu(value IComponent) {
    ScrollBar_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TScrollBar) Position() int32 {
    return ScrollBar_GetPosition(s.instance)
}

func (s *TScrollBar) SetPosition(value int32) {
    ScrollBar_SetPosition(s.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TScrollBar) ShowHint() bool {
    return ScrollBar_GetShowHint(s.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TScrollBar) SetShowHint(value bool) {
    ScrollBar_SetShowHint(s.instance, value)
}

func (s *TScrollBar) SmallChange() TScrollBarInc {
    return ScrollBar_GetSmallChange(s.instance)
}

func (s *TScrollBar) SetSmallChange(value TScrollBarInc) {
    ScrollBar_SetSmallChange(s.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TScrollBar) TabOrder() TTabOrder {
    return ScrollBar_GetTabOrder(s.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TScrollBar) SetTabOrder(value TTabOrder) {
    ScrollBar_SetTabOrder(s.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TScrollBar) TabStop() bool {
    return ScrollBar_GetTabStop(s.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TScrollBar) SetTabStop(value bool) {
    ScrollBar_SetTabStop(s.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TScrollBar) Visible() bool {
    return ScrollBar_GetVisible(s.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TScrollBar) SetVisible(value bool) {
    ScrollBar_SetVisible(s.instance, value)
}

// CN: 获取样式元素。
// EN: Get Style element.
func (s *TScrollBar) StyleElements() TStyleElements {
    return ScrollBar_GetStyleElements(s.instance)
}

// CN: 设置样式元素。
// EN: Set Style element.
func (s *TScrollBar) SetStyleElements(value TStyleElements) {
    ScrollBar_SetStyleElements(s.instance, value)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (s *TScrollBar) SetOnContextPopup(fn TContextPopupEvent) {
    ScrollBar_SetOnContextPopup(s.instance, fn)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (s *TScrollBar) SetOnChange(fn TNotifyEvent) {
    ScrollBar_SetOnChange(s.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (s *TScrollBar) SetOnDragDrop(fn TDragDropEvent) {
    ScrollBar_SetOnDragDrop(s.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (s *TScrollBar) SetOnDragOver(fn TDragOverEvent) {
    ScrollBar_SetOnDragOver(s.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (s *TScrollBar) SetOnEndDock(fn TEndDragEvent) {
    ScrollBar_SetOnEndDock(s.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (s *TScrollBar) SetOnEndDrag(fn TEndDragEvent) {
    ScrollBar_SetOnEndDrag(s.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (s *TScrollBar) SetOnEnter(fn TNotifyEvent) {
    ScrollBar_SetOnEnter(s.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (s *TScrollBar) SetOnExit(fn TNotifyEvent) {
    ScrollBar_SetOnExit(s.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (s *TScrollBar) SetOnKeyDown(fn TKeyEvent) {
    ScrollBar_SetOnKeyDown(s.instance, fn)
}

func (s *TScrollBar) SetOnKeyPress(fn TKeyPressEvent) {
    ScrollBar_SetOnKeyPress(s.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (s *TScrollBar) SetOnKeyUp(fn TKeyEvent) {
    ScrollBar_SetOnKeyUp(s.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (s *TScrollBar) SetOnMouseEnter(fn TNotifyEvent) {
    ScrollBar_SetOnMouseEnter(s.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (s *TScrollBar) SetOnMouseLeave(fn TNotifyEvent) {
    ScrollBar_SetOnMouseLeave(s.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (s *TScrollBar) SetOnStartDock(fn TStartDockEvent) {
    ScrollBar_SetOnStartDock(s.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (s *TScrollBar) DockClientCount() int32 {
    return ScrollBar_GetDockClientCount(s.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TScrollBar) DockSite() bool {
    return ScrollBar_GetDockSite(s.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TScrollBar) SetDockSite(value bool) {
    ScrollBar_SetDockSite(s.instance, value)
}

// CN: 获取禁用对齐。
// EN: .
func (s *TScrollBar) AlignDisabled() bool {
    return ScrollBar_GetAlignDisabled(s.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TScrollBar) MouseInClient() bool {
    return ScrollBar_GetMouseInClient(s.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TScrollBar) VisibleDockClientCount() int32 {
    return ScrollBar_GetVisibleDockClientCount(s.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TScrollBar) Brush() *TBrush {
    return AsBrush(ScrollBar_GetBrush(s.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TScrollBar) ControlCount() int32 {
    return ScrollBar_GetControlCount(s.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TScrollBar) Handle() HWND {
    return ScrollBar_GetHandle(s.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TScrollBar) ParentWindow() HWND {
    return ScrollBar_GetParentWindow(s.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TScrollBar) SetParentWindow(value HWND) {
    ScrollBar_SetParentWindow(s.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (s *TScrollBar) UseDockManager() bool {
    return ScrollBar_GetUseDockManager(s.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (s *TScrollBar) SetUseDockManager(value bool) {
    ScrollBar_SetUseDockManager(s.instance, value)
}

func (s *TScrollBar) Action() *TAction {
    return AsAction(ScrollBar_GetAction(s.instance))
}

func (s *TScrollBar) SetAction(value IComponent) {
    ScrollBar_SetAction(s.instance, CheckPtr(value))
}

func (s *TScrollBar) BoundsRect() TRect {
    return ScrollBar_GetBoundsRect(s.instance)
}

func (s *TScrollBar) SetBoundsRect(value TRect) {
    ScrollBar_SetBoundsRect(s.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (s *TScrollBar) ClientHeight() int32 {
    return ScrollBar_GetClientHeight(s.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (s *TScrollBar) SetClientHeight(value int32) {
    ScrollBar_SetClientHeight(s.instance, value)
}

func (s *TScrollBar) ClientOrigin() TPoint {
    return ScrollBar_GetClientOrigin(s.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TScrollBar) ClientRect() TRect {
    return ScrollBar_GetClientRect(s.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TScrollBar) ClientWidth() int32 {
    return ScrollBar_GetClientWidth(s.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TScrollBar) SetClientWidth(value int32) {
    ScrollBar_SetClientWidth(s.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (s *TScrollBar) ControlState() TControlState {
    return ScrollBar_GetControlState(s.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (s *TScrollBar) SetControlState(value TControlState) {
    ScrollBar_SetControlState(s.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (s *TScrollBar) ControlStyle() TControlStyle {
    return ScrollBar_GetControlStyle(s.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (s *TScrollBar) SetControlStyle(value TControlStyle) {
    ScrollBar_SetControlStyle(s.instance, value)
}

func (s *TScrollBar) ExplicitLeft() int32 {
    return ScrollBar_GetExplicitLeft(s.instance)
}

func (s *TScrollBar) ExplicitTop() int32 {
    return ScrollBar_GetExplicitTop(s.instance)
}

func (s *TScrollBar) ExplicitWidth() int32 {
    return ScrollBar_GetExplicitWidth(s.instance)
}

func (s *TScrollBar) ExplicitHeight() int32 {
    return ScrollBar_GetExplicitHeight(s.instance)
}

func (s *TScrollBar) Floating() bool {
    return ScrollBar_GetFloating(s.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TScrollBar) Parent() *TWinControl {
    return AsWinControl(ScrollBar_GetParent(s.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TScrollBar) SetParent(value IWinControl) {
    ScrollBar_SetParent(s.instance, CheckPtr(value))
}

func (s *TScrollBar) SetOnGesture(fn TGestureEvent) {
    ScrollBar_SetOnGesture(s.instance, fn)
}

// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TScrollBar) AlignWithMargins() bool {
    return ScrollBar_GetAlignWithMargins(s.instance)
}

// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TScrollBar) SetAlignWithMargins(value bool) {
    ScrollBar_SetAlignWithMargins(s.instance, value)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (s *TScrollBar) Left() int32 {
    return ScrollBar_GetLeft(s.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (s *TScrollBar) SetLeft(value int32) {
    ScrollBar_SetLeft(s.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TScrollBar) Top() int32 {
    return ScrollBar_GetTop(s.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TScrollBar) SetTop(value int32) {
    ScrollBar_SetTop(s.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (s *TScrollBar) Width() int32 {
    return ScrollBar_GetWidth(s.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (s *TScrollBar) SetWidth(value int32) {
    ScrollBar_SetWidth(s.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (s *TScrollBar) Height() int32 {
    return ScrollBar_GetHeight(s.instance)
}

// CN: 设置高度。
// EN: Set height.
func (s *TScrollBar) SetHeight(value int32) {
    ScrollBar_SetHeight(s.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TScrollBar) Cursor() TCursor {
    return ScrollBar_GetCursor(s.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TScrollBar) SetCursor(value TCursor) {
    ScrollBar_SetCursor(s.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TScrollBar) Hint() string {
    return ScrollBar_GetHint(s.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TScrollBar) SetHint(value string) {
    ScrollBar_SetHint(s.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TScrollBar) Margins() *TMargins {
    return AsMargins(ScrollBar_GetMargins(s.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TScrollBar) SetMargins(value *TMargins) {
    ScrollBar_SetMargins(s.instance, CheckPtr(value))
}

// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TScrollBar) CustomHint() *TCustomHint {
    return AsCustomHint(ScrollBar_GetCustomHint(s.instance))
}

// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TScrollBar) SetCustomHint(value IComponent) {
    ScrollBar_SetCustomHint(s.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TScrollBar) ComponentCount() int32 {
    return ScrollBar_GetComponentCount(s.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (s *TScrollBar) ComponentIndex() int32 {
    return ScrollBar_GetComponentIndex(s.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (s *TScrollBar) SetComponentIndex(value int32) {
    ScrollBar_SetComponentIndex(s.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TScrollBar) Owner() *TComponent {
    return AsComponent(ScrollBar_GetOwner(s.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (s *TScrollBar) Name() string {
    return ScrollBar_GetName(s.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (s *TScrollBar) SetName(value string) {
    ScrollBar_SetName(s.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TScrollBar) Tag() int {
    return ScrollBar_GetTag(s.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TScrollBar) SetTag(value int) {
    ScrollBar_SetTag(s.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (s *TScrollBar) DockClients(Index int32) *TControl {
    return AsControl(ScrollBar_GetDockClients(s.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (s *TScrollBar) Controls(Index int32) *TControl {
    return AsControl(ScrollBar_GetControls(s.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TScrollBar) Components(AIndex int32) *TComponent {
    return AsComponent(ScrollBar_GetComponents(s.instance, AIndex))
}

