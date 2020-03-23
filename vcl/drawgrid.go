
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

type TDrawGrid struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDrawGrid(owner IComponent) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = DrawGrid_Create(CheckPtr(owner))
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsDrawGrid(obj interface{}) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance, d.ptr = getInstance(obj)
    if d.instance == 0 { return nil }
    return d
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsDrawGrid.
func DrawGridFromInst(inst uintptr) *TDrawGrid {
    return AsDrawGrid(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsDrawGrid.
func DrawGridFromObj(obj IObject) *TDrawGrid {
    return AsDrawGrid(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsDrawGrid.
func DrawGridFromUnsafePointer(ptr unsafe.Pointer) *TDrawGrid {
    return AsDrawGrid(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (d *TDrawGrid) Free() {
    if d.instance != 0 {
        DrawGrid_Free(d.instance)
        d.instance, d.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDrawGrid) Instance() uintptr {
    return d.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDrawGrid) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDrawGrid) IsValid() bool {
    return d.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (d *TDrawGrid) Is() TIs {
    return TIs(d.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (d *TDrawGrid) As() TAs {
//    return TAs(d.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDrawGridClass() TClass {
    return DrawGrid_StaticClassType()
}

func (d *TDrawGrid) CellRect(ACol int32, ARow int32) TRect {
    return DrawGrid_CellRect(d.instance, ACol , ARow)
}

func (d *TDrawGrid) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    DrawGrid_MouseToCell(d.instance, X , Y , ACol , ARow)
}

func (d *TDrawGrid) MouseCoord(X int32, Y int32) TGridCoord {
    return DrawGrid_MouseCoord(d.instance, X , Y)
}

// CN: 是否可以获得焦点。
// EN: .
func (d *TDrawGrid) CanFocus() bool {
    return DrawGrid_CanFocus(d.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (d *TDrawGrid) ContainsControl(Control IControl) bool {
    return DrawGrid_ContainsControl(d.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (d *TDrawGrid) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(DrawGrid_ControlAtPos(d.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (d *TDrawGrid) DisableAlign() {
    DrawGrid_DisableAlign(d.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (d *TDrawGrid) EnableAlign() {
    DrawGrid_EnableAlign(d.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (d *TDrawGrid) FindChildControl(ControlName string) *TControl {
    return AsControl(DrawGrid_FindChildControl(d.instance, ControlName))
}

func (d *TDrawGrid) FlipChildren(AllLevels bool) {
    DrawGrid_FlipChildren(d.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (d *TDrawGrid) Focused() bool {
    return DrawGrid_Focused(d.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (d *TDrawGrid) HandleAllocated() bool {
    return DrawGrid_HandleAllocated(d.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (d *TDrawGrid) InsertControl(AControl IControl) {
    DrawGrid_InsertControl(d.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (d *TDrawGrid) Invalidate() {
    DrawGrid_Invalidate(d.instance)
}

// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (d *TDrawGrid) PaintTo(DC HDC, X int32, Y int32) {
    DrawGrid_PaintTo(d.instance, DC , X , Y)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (d *TDrawGrid) RemoveControl(AControl IControl) {
    DrawGrid_RemoveControl(d.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (d *TDrawGrid) Realign() {
    DrawGrid_Realign(d.instance)
}

// CN: 重绘。
// EN: Repaint.
func (d *TDrawGrid) Repaint() {
    DrawGrid_Repaint(d.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (d *TDrawGrid) ScaleBy(M int32, D int32) {
    DrawGrid_ScaleBy(d.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (d *TDrawGrid) ScrollBy(DeltaX int32, DeltaY int32) {
    DrawGrid_ScrollBy(d.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (d *TDrawGrid) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DrawGrid_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (d *TDrawGrid) SetFocus() {
    DrawGrid_SetFocus(d.instance)
}

// CN: 控件更新。
// EN: Update.
func (d *TDrawGrid) Update() {
    DrawGrid_Update(d.instance)
}

// CN: 更新控件状态。
// EN: Update control status.
func (d *TDrawGrid) UpdateControlState() {
    DrawGrid_UpdateControlState(d.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (d *TDrawGrid) BringToFront() {
    DrawGrid_BringToFront(d.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (d *TDrawGrid) ClientToScreen(Point TPoint) TPoint {
    return DrawGrid_ClientToScreen(d.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (d *TDrawGrid) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return DrawGrid_ClientToParent(d.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (d *TDrawGrid) Dragging() bool {
    return DrawGrid_Dragging(d.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (d *TDrawGrid) HasParent() bool {
    return DrawGrid_HasParent(d.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (d *TDrawGrid) Hide() {
    DrawGrid_Hide(d.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (d *TDrawGrid) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DrawGrid_Perform(d.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (d *TDrawGrid) Refresh() {
    DrawGrid_Refresh(d.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (d *TDrawGrid) ScreenToClient(Point TPoint) TPoint {
    return DrawGrid_ScreenToClient(d.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (d *TDrawGrid) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return DrawGrid_ParentToClient(d.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (d *TDrawGrid) SendToBack() {
    DrawGrid_SendToBack(d.instance)
}

// CN: 显示控件。
// EN: Show control.
func (d *TDrawGrid) Show() {
    DrawGrid_Show(d.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (d *TDrawGrid) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return DrawGrid_GetTextBuf(d.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (d *TDrawGrid) GetTextLen() int32 {
    return DrawGrid_GetTextLen(d.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (d *TDrawGrid) SetTextBuf(Buffer string) {
    DrawGrid_SetTextBuf(d.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (d *TDrawGrid) FindComponent(AName string) *TComponent {
    return AsComponent(DrawGrid_FindComponent(d.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (d *TDrawGrid) GetNamePath() string {
    return DrawGrid_GetNamePath(d.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (d *TDrawGrid) Assign(Source IObject) {
    DrawGrid_Assign(d.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDrawGrid) DisposeOf() {
    DrawGrid_DisposeOf(d.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDrawGrid) ClassType() TClass {
    return DrawGrid_ClassType(d.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDrawGrid) ClassName() string {
    return DrawGrid_ClassName(d.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDrawGrid) InstanceSize() int32 {
    return DrawGrid_InstanceSize(d.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDrawGrid) InheritsFrom(AClass TClass) bool {
    return DrawGrid_InheritsFrom(d.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDrawGrid) Equals(Obj IObject) bool {
    return DrawGrid_Equals(d.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDrawGrid) GetHashCode() int32 {
    return DrawGrid_GetHashCode(d.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (d *TDrawGrid) ToString() string {
    return DrawGrid_ToString(d.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (d *TDrawGrid) Align() TAlign {
    return DrawGrid_GetAlign(d.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (d *TDrawGrid) SetAlign(value TAlign) {
    DrawGrid_SetAlign(d.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (d *TDrawGrid) Anchors() TAnchors {
    return DrawGrid_GetAnchors(d.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (d *TDrawGrid) SetAnchors(value TAnchors) {
    DrawGrid_SetAnchors(d.instance, value)
}

func (d *TDrawGrid) BevelEdges() TBevelEdges {
    return DrawGrid_GetBevelEdges(d.instance)
}

func (d *TDrawGrid) SetBevelEdges(value TBevelEdges) {
    DrawGrid_SetBevelEdges(d.instance, value)
}

func (d *TDrawGrid) BevelInner() TBevelCut {
    return DrawGrid_GetBevelInner(d.instance)
}

func (d *TDrawGrid) SetBevelInner(value TBevelCut) {
    DrawGrid_SetBevelInner(d.instance, value)
}

func (d *TDrawGrid) BevelKind() TBevelKind {
    return DrawGrid_GetBevelKind(d.instance)
}

func (d *TDrawGrid) SetBevelKind(value TBevelKind) {
    DrawGrid_SetBevelKind(d.instance, value)
}

func (d *TDrawGrid) BevelOuter() TBevelCut {
    return DrawGrid_GetBevelOuter(d.instance)
}

func (d *TDrawGrid) SetBevelOuter(value TBevelCut) {
    DrawGrid_SetBevelOuter(d.instance, value)
}

func (d *TDrawGrid) BiDiMode() TBiDiMode {
    return DrawGrid_GetBiDiMode(d.instance)
}

func (d *TDrawGrid) SetBiDiMode(value TBiDiMode) {
    DrawGrid_SetBiDiMode(d.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (d *TDrawGrid) BorderStyle() TBorderStyle {
    return DrawGrid_GetBorderStyle(d.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (d *TDrawGrid) SetBorderStyle(value TBorderStyle) {
    DrawGrid_SetBorderStyle(d.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (d *TDrawGrid) Color() TColor {
    return DrawGrid_GetColor(d.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (d *TDrawGrid) SetColor(value TColor) {
    DrawGrid_SetColor(d.instance, value)
}

func (d *TDrawGrid) ColCount() int32 {
    return DrawGrid_GetColCount(d.instance)
}

func (d *TDrawGrid) SetColCount(value int32) {
    DrawGrid_SetColCount(d.instance, value)
}

func (d *TDrawGrid) Constraints() *TSizeConstraints {
    return AsSizeConstraints(DrawGrid_GetConstraints(d.instance))
}

func (d *TDrawGrid) SetConstraints(value *TSizeConstraints) {
    DrawGrid_SetConstraints(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) Ctl3D() bool {
    return DrawGrid_GetCtl3D(d.instance)
}

func (d *TDrawGrid) SetCtl3D(value bool) {
    DrawGrid_SetCtl3D(d.instance, value)
}

func (d *TDrawGrid) DefaultColWidth() int32 {
    return DrawGrid_GetDefaultColWidth(d.instance)
}

func (d *TDrawGrid) SetDefaultColWidth(value int32) {
    DrawGrid_SetDefaultColWidth(d.instance, value)
}

func (d *TDrawGrid) DefaultRowHeight() int32 {
    return DrawGrid_GetDefaultRowHeight(d.instance)
}

func (d *TDrawGrid) SetDefaultRowHeight(value int32) {
    DrawGrid_SetDefaultRowHeight(d.instance, value)
}

func (d *TDrawGrid) DefaultDrawing() bool {
    return DrawGrid_GetDefaultDrawing(d.instance)
}

func (d *TDrawGrid) SetDefaultDrawing(value bool) {
    DrawGrid_SetDefaultDrawing(d.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (d *TDrawGrid) DoubleBuffered() bool {
    return DrawGrid_GetDoubleBuffered(d.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (d *TDrawGrid) SetDoubleBuffered(value bool) {
    DrawGrid_SetDoubleBuffered(d.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (d *TDrawGrid) DragCursor() TCursor {
    return DrawGrid_GetDragCursor(d.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (d *TDrawGrid) SetDragCursor(value TCursor) {
    DrawGrid_SetDragCursor(d.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (d *TDrawGrid) DragKind() TDragKind {
    return DrawGrid_GetDragKind(d.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (d *TDrawGrid) SetDragKind(value TDragKind) {
    DrawGrid_SetDragKind(d.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (d *TDrawGrid) DragMode() TDragMode {
    return DrawGrid_GetDragMode(d.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (d *TDrawGrid) SetDragMode(value TDragMode) {
    DrawGrid_SetDragMode(d.instance, value)
}

func (d *TDrawGrid) DrawingStyle() TGridDrawingStyle {
    return DrawGrid_GetDrawingStyle(d.instance)
}

func (d *TDrawGrid) SetDrawingStyle(value TGridDrawingStyle) {
    DrawGrid_SetDrawingStyle(d.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (d *TDrawGrid) Enabled() bool {
    return DrawGrid_GetEnabled(d.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (d *TDrawGrid) SetEnabled(value bool) {
    DrawGrid_SetEnabled(d.instance, value)
}

func (d *TDrawGrid) FixedColor() TColor {
    return DrawGrid_GetFixedColor(d.instance)
}

func (d *TDrawGrid) SetFixedColor(value TColor) {
    DrawGrid_SetFixedColor(d.instance, value)
}

func (d *TDrawGrid) FixedCols() int32 {
    return DrawGrid_GetFixedCols(d.instance)
}

func (d *TDrawGrid) SetFixedCols(value int32) {
    DrawGrid_SetFixedCols(d.instance, value)
}

func (d *TDrawGrid) RowCount() int32 {
    return DrawGrid_GetRowCount(d.instance)
}

func (d *TDrawGrid) SetRowCount(value int32) {
    DrawGrid_SetRowCount(d.instance, value)
}

func (d *TDrawGrid) FixedRows() int32 {
    return DrawGrid_GetFixedRows(d.instance)
}

func (d *TDrawGrid) SetFixedRows(value int32) {
    DrawGrid_SetFixedRows(d.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (d *TDrawGrid) Font() *TFont {
    return AsFont(DrawGrid_GetFont(d.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (d *TDrawGrid) SetFont(value *TFont) {
    DrawGrid_SetFont(d.instance, CheckPtr(value))
}

// CN: 获取渐变结束颜色, 仅VCL。
// EN: .
func (d *TDrawGrid) GradientEndColor() TColor {
    return DrawGrid_GetGradientEndColor(d.instance)
}

// CN: 设置渐变结束颜色, 仅VCL。
// EN: .
func (d *TDrawGrid) SetGradientEndColor(value TColor) {
    DrawGrid_SetGradientEndColor(d.instance, value)
}

// CN: 获取渐变起始颜色，仅VCL。
// EN: .
func (d *TDrawGrid) GradientStartColor() TColor {
    return DrawGrid_GetGradientStartColor(d.instance)
}

// CN: 设置渐变起始颜色，仅VCL。
// EN: .
func (d *TDrawGrid) SetGradientStartColor(value TColor) {
    DrawGrid_SetGradientStartColor(d.instance, value)
}

func (d *TDrawGrid) GridLineWidth() int32 {
    return DrawGrid_GetGridLineWidth(d.instance)
}

func (d *TDrawGrid) SetGridLineWidth(value int32) {
    DrawGrid_SetGridLineWidth(d.instance, value)
}

func (d *TDrawGrid) Options() TGridOptions {
    return DrawGrid_GetOptions(d.instance)
}

func (d *TDrawGrid) SetOptions(value TGridOptions) {
    DrawGrid_SetOptions(d.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (d *TDrawGrid) ParentColor() bool {
    return DrawGrid_GetParentColor(d.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (d *TDrawGrid) SetParentColor(value bool) {
    DrawGrid_SetParentColor(d.instance, value)
}

func (d *TDrawGrid) ParentCtl3D() bool {
    return DrawGrid_GetParentCtl3D(d.instance)
}

func (d *TDrawGrid) SetParentCtl3D(value bool) {
    DrawGrid_SetParentCtl3D(d.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (d *TDrawGrid) ParentDoubleBuffered() bool {
    return DrawGrid_GetParentDoubleBuffered(d.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (d *TDrawGrid) SetParentDoubleBuffered(value bool) {
    DrawGrid_SetParentDoubleBuffered(d.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (d *TDrawGrid) ParentFont() bool {
    return DrawGrid_GetParentFont(d.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (d *TDrawGrid) SetParentFont(value bool) {
    DrawGrid_SetParentFont(d.instance, value)
}

func (d *TDrawGrid) ParentShowHint() bool {
    return DrawGrid_GetParentShowHint(d.instance)
}

func (d *TDrawGrid) SetParentShowHint(value bool) {
    DrawGrid_SetParentShowHint(d.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (d *TDrawGrid) PopupMenu() *TPopupMenu {
    return AsPopupMenu(DrawGrid_GetPopupMenu(d.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (d *TDrawGrid) SetPopupMenu(value IComponent) {
    DrawGrid_SetPopupMenu(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) ScrollBars() TScrollStyle {
    return DrawGrid_GetScrollBars(d.instance)
}

func (d *TDrawGrid) SetScrollBars(value TScrollStyle) {
    DrawGrid_SetScrollBars(d.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (d *TDrawGrid) ShowHint() bool {
    return DrawGrid_GetShowHint(d.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (d *TDrawGrid) SetShowHint(value bool) {
    DrawGrid_SetShowHint(d.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (d *TDrawGrid) TabOrder() TTabOrder {
    return DrawGrid_GetTabOrder(d.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (d *TDrawGrid) SetTabOrder(value TTabOrder) {
    DrawGrid_SetTabOrder(d.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (d *TDrawGrid) Visible() bool {
    return DrawGrid_GetVisible(d.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (d *TDrawGrid) SetVisible(value bool) {
    DrawGrid_SetVisible(d.instance, value)
}

// CN: 获取样式元素。
// EN: Get Style element.
func (d *TDrawGrid) StyleElements() TStyleElements {
    return DrawGrid_GetStyleElements(d.instance)
}

// CN: 设置样式元素。
// EN: Set Style element.
func (d *TDrawGrid) SetStyleElements(value TStyleElements) {
    DrawGrid_SetStyleElements(d.instance, value)
}

func (d *TDrawGrid) VisibleColCount() int32 {
    return DrawGrid_GetVisibleColCount(d.instance)
}

func (d *TDrawGrid) VisibleRowCount() int32 {
    return DrawGrid_GetVisibleRowCount(d.instance)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (d *TDrawGrid) SetOnClick(fn TNotifyEvent) {
    DrawGrid_SetOnClick(d.instance, fn)
}

func (d *TDrawGrid) SetOnColumnMoved(fn TMovedEvent) {
    DrawGrid_SetOnColumnMoved(d.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (d *TDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
    DrawGrid_SetOnContextPopup(d.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (d *TDrawGrid) SetOnDblClick(fn TNotifyEvent) {
    DrawGrid_SetOnDblClick(d.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (d *TDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
    DrawGrid_SetOnDragDrop(d.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (d *TDrawGrid) SetOnDragOver(fn TDragOverEvent) {
    DrawGrid_SetOnDragOver(d.instance, fn)
}

func (d *TDrawGrid) SetOnDrawCell(fn TDrawCellEvent) {
    DrawGrid_SetOnDrawCell(d.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (d *TDrawGrid) SetOnEndDock(fn TEndDragEvent) {
    DrawGrid_SetOnEndDock(d.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (d *TDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
    DrawGrid_SetOnEndDrag(d.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (d *TDrawGrid) SetOnEnter(fn TNotifyEvent) {
    DrawGrid_SetOnEnter(d.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (d *TDrawGrid) SetOnExit(fn TNotifyEvent) {
    DrawGrid_SetOnExit(d.instance, fn)
}

func (d *TDrawGrid) SetOnFixedCellClick(fn TFixedCellClickEvent) {
    DrawGrid_SetOnFixedCellClick(d.instance, fn)
}

func (d *TDrawGrid) SetOnGesture(fn TGestureEvent) {
    DrawGrid_SetOnGesture(d.instance, fn)
}

func (d *TDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
    DrawGrid_SetOnGetEditMask(d.instance, fn)
}

func (d *TDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
    DrawGrid_SetOnGetEditText(d.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (d *TDrawGrid) SetOnKeyDown(fn TKeyEvent) {
    DrawGrid_SetOnKeyDown(d.instance, fn)
}

func (d *TDrawGrid) SetOnKeyPress(fn TKeyPressEvent) {
    DrawGrid_SetOnKeyPress(d.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (d *TDrawGrid) SetOnKeyUp(fn TKeyEvent) {
    DrawGrid_SetOnKeyUp(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseActivate(fn TMouseActivateEvent) {
    DrawGrid_SetOnMouseActivate(d.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (d *TDrawGrid) SetOnMouseDown(fn TMouseEvent) {
    DrawGrid_SetOnMouseDown(d.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (d *TDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
    DrawGrid_SetOnMouseEnter(d.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (d *TDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
    DrawGrid_SetOnMouseLeave(d.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (d *TDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
    DrawGrid_SetOnMouseMove(d.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (d *TDrawGrid) SetOnMouseUp(fn TMouseEvent) {
    DrawGrid_SetOnMouseUp(d.instance, fn)
}

// CN: 设置鼠标滚轮按下事件。
// EN: .
func (d *TDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    DrawGrid_SetOnMouseWheelDown(d.instance, fn)
}

// CN: 设置鼠标滚轮抬起事件。
// EN: .
func (d *TDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    DrawGrid_SetOnMouseWheelUp(d.instance, fn)
}

func (d *TDrawGrid) SetOnRowMoved(fn TMovedEvent) {
    DrawGrid_SetOnRowMoved(d.instance, fn)
}

func (d *TDrawGrid) SetOnSelectCell(fn TSelectCellEvent) {
    DrawGrid_SetOnSelectCell(d.instance, fn)
}

func (d *TDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
    DrawGrid_SetOnSetEditText(d.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (d *TDrawGrid) SetOnStartDock(fn TStartDockEvent) {
    DrawGrid_SetOnStartDock(d.instance, fn)
}

func (d *TDrawGrid) SetOnTopLeftChanged(fn TNotifyEvent) {
    DrawGrid_SetOnTopLeftChanged(d.instance, fn)
}

// CN: 获取画布。
// EN: .
func (d *TDrawGrid) Canvas() *TCanvas {
    return AsCanvas(DrawGrid_GetCanvas(d.instance))
}

func (d *TDrawGrid) Col() int32 {
    return DrawGrid_GetCol(d.instance)
}

func (d *TDrawGrid) SetCol(value int32) {
    DrawGrid_SetCol(d.instance, value)
}

func (d *TDrawGrid) EditorMode() bool {
    return DrawGrid_GetEditorMode(d.instance)
}

func (d *TDrawGrid) SetEditorMode(value bool) {
    DrawGrid_SetEditorMode(d.instance, value)
}

func (d *TDrawGrid) GridHeight() int32 {
    return DrawGrid_GetGridHeight(d.instance)
}

func (d *TDrawGrid) GridWidth() int32 {
    return DrawGrid_GetGridWidth(d.instance)
}

func (d *TDrawGrid) LeftCol() int32 {
    return DrawGrid_GetLeftCol(d.instance)
}

func (d *TDrawGrid) SetLeftCol(value int32) {
    DrawGrid_SetLeftCol(d.instance, value)
}

func (d *TDrawGrid) Selection() TGridRect {
    return DrawGrid_GetSelection(d.instance)
}

func (d *TDrawGrid) SetSelection(value TGridRect) {
    DrawGrid_SetSelection(d.instance, value)
}

func (d *TDrawGrid) Row() int32 {
    return DrawGrid_GetRow(d.instance)
}

func (d *TDrawGrid) SetRow(value int32) {
    DrawGrid_SetRow(d.instance, value)
}

func (d *TDrawGrid) TopRow() int32 {
    return DrawGrid_GetTopRow(d.instance)
}

func (d *TDrawGrid) SetTopRow(value int32) {
    DrawGrid_SetTopRow(d.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (d *TDrawGrid) TabStop() bool {
    return DrawGrid_GetTabStop(d.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (d *TDrawGrid) SetTabStop(value bool) {
    DrawGrid_SetTabStop(d.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (d *TDrawGrid) DockClientCount() int32 {
    return DrawGrid_GetDockClientCount(d.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (d *TDrawGrid) DockSite() bool {
    return DrawGrid_GetDockSite(d.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (d *TDrawGrid) SetDockSite(value bool) {
    DrawGrid_SetDockSite(d.instance, value)
}

// CN: 获取禁用对齐。
// EN: .
func (d *TDrawGrid) AlignDisabled() bool {
    return DrawGrid_GetAlignDisabled(d.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (d *TDrawGrid) MouseInClient() bool {
    return DrawGrid_GetMouseInClient(d.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (d *TDrawGrid) VisibleDockClientCount() int32 {
    return DrawGrid_GetVisibleDockClientCount(d.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (d *TDrawGrid) Brush() *TBrush {
    return AsBrush(DrawGrid_GetBrush(d.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (d *TDrawGrid) ControlCount() int32 {
    return DrawGrid_GetControlCount(d.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (d *TDrawGrid) Handle() HWND {
    return DrawGrid_GetHandle(d.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (d *TDrawGrid) ParentWindow() HWND {
    return DrawGrid_GetParentWindow(d.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (d *TDrawGrid) SetParentWindow(value HWND) {
    DrawGrid_SetParentWindow(d.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (d *TDrawGrid) UseDockManager() bool {
    return DrawGrid_GetUseDockManager(d.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (d *TDrawGrid) SetUseDockManager(value bool) {
    DrawGrid_SetUseDockManager(d.instance, value)
}

func (d *TDrawGrid) Action() *TAction {
    return AsAction(DrawGrid_GetAction(d.instance))
}

func (d *TDrawGrid) SetAction(value IComponent) {
    DrawGrid_SetAction(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) BoundsRect() TRect {
    return DrawGrid_GetBoundsRect(d.instance)
}

func (d *TDrawGrid) SetBoundsRect(value TRect) {
    DrawGrid_SetBoundsRect(d.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (d *TDrawGrid) ClientHeight() int32 {
    return DrawGrid_GetClientHeight(d.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (d *TDrawGrid) SetClientHeight(value int32) {
    DrawGrid_SetClientHeight(d.instance, value)
}

func (d *TDrawGrid) ClientOrigin() TPoint {
    return DrawGrid_GetClientOrigin(d.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (d *TDrawGrid) ClientRect() TRect {
    return DrawGrid_GetClientRect(d.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (d *TDrawGrid) ClientWidth() int32 {
    return DrawGrid_GetClientWidth(d.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (d *TDrawGrid) SetClientWidth(value int32) {
    DrawGrid_SetClientWidth(d.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (d *TDrawGrid) ControlState() TControlState {
    return DrawGrid_GetControlState(d.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (d *TDrawGrid) SetControlState(value TControlState) {
    DrawGrid_SetControlState(d.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (d *TDrawGrid) ControlStyle() TControlStyle {
    return DrawGrid_GetControlStyle(d.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (d *TDrawGrid) SetControlStyle(value TControlStyle) {
    DrawGrid_SetControlStyle(d.instance, value)
}

func (d *TDrawGrid) ExplicitLeft() int32 {
    return DrawGrid_GetExplicitLeft(d.instance)
}

func (d *TDrawGrid) ExplicitTop() int32 {
    return DrawGrid_GetExplicitTop(d.instance)
}

func (d *TDrawGrid) ExplicitWidth() int32 {
    return DrawGrid_GetExplicitWidth(d.instance)
}

func (d *TDrawGrid) ExplicitHeight() int32 {
    return DrawGrid_GetExplicitHeight(d.instance)
}

func (d *TDrawGrid) Floating() bool {
    return DrawGrid_GetFloating(d.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (d *TDrawGrid) Parent() *TWinControl {
    return AsWinControl(DrawGrid_GetParent(d.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (d *TDrawGrid) SetParent(value IWinControl) {
    DrawGrid_SetParent(d.instance, CheckPtr(value))
}

// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (d *TDrawGrid) AlignWithMargins() bool {
    return DrawGrid_GetAlignWithMargins(d.instance)
}

// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (d *TDrawGrid) SetAlignWithMargins(value bool) {
    DrawGrid_SetAlignWithMargins(d.instance, value)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (d *TDrawGrid) Left() int32 {
    return DrawGrid_GetLeft(d.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (d *TDrawGrid) SetLeft(value int32) {
    DrawGrid_SetLeft(d.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (d *TDrawGrid) Top() int32 {
    return DrawGrid_GetTop(d.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (d *TDrawGrid) SetTop(value int32) {
    DrawGrid_SetTop(d.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (d *TDrawGrid) Width() int32 {
    return DrawGrid_GetWidth(d.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (d *TDrawGrid) SetWidth(value int32) {
    DrawGrid_SetWidth(d.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (d *TDrawGrid) Height() int32 {
    return DrawGrid_GetHeight(d.instance)
}

// CN: 设置高度。
// EN: Set height.
func (d *TDrawGrid) SetHeight(value int32) {
    DrawGrid_SetHeight(d.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (d *TDrawGrid) Cursor() TCursor {
    return DrawGrid_GetCursor(d.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (d *TDrawGrid) SetCursor(value TCursor) {
    DrawGrid_SetCursor(d.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (d *TDrawGrid) Hint() string {
    return DrawGrid_GetHint(d.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (d *TDrawGrid) SetHint(value string) {
    DrawGrid_SetHint(d.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (d *TDrawGrid) Margins() *TMargins {
    return AsMargins(DrawGrid_GetMargins(d.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (d *TDrawGrid) SetMargins(value *TMargins) {
    DrawGrid_SetMargins(d.instance, CheckPtr(value))
}

// CN: 获取自定义提示。
// EN: Get custom hint.
func (d *TDrawGrid) CustomHint() *TCustomHint {
    return AsCustomHint(DrawGrid_GetCustomHint(d.instance))
}

// CN: 设置自定义提示。
// EN: Set custom hint.
func (d *TDrawGrid) SetCustomHint(value IComponent) {
    DrawGrid_SetCustomHint(d.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (d *TDrawGrid) ComponentCount() int32 {
    return DrawGrid_GetComponentCount(d.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (d *TDrawGrid) ComponentIndex() int32 {
    return DrawGrid_GetComponentIndex(d.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (d *TDrawGrid) SetComponentIndex(value int32) {
    DrawGrid_SetComponentIndex(d.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (d *TDrawGrid) Owner() *TComponent {
    return AsComponent(DrawGrid_GetOwner(d.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (d *TDrawGrid) Name() string {
    return DrawGrid_GetName(d.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (d *TDrawGrid) SetName(value string) {
    DrawGrid_SetName(d.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (d *TDrawGrid) Tag() int {
    return DrawGrid_GetTag(d.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (d *TDrawGrid) SetTag(value int) {
    DrawGrid_SetTag(d.instance, value)
}

func (d *TDrawGrid) ColWidths(Index int32) int32 {
    return DrawGrid_GetColWidths(d.instance, Index)
}

func (d *TDrawGrid) SetColWidths(Index int32, value int32) {
    DrawGrid_SetColWidths(d.instance, Index, value)
}

func (d *TDrawGrid) RowHeights(Index int32) int32 {
    return DrawGrid_GetRowHeights(d.instance, Index)
}

func (d *TDrawGrid) SetRowHeights(Index int32, value int32) {
    DrawGrid_SetRowHeights(d.instance, Index, value)
}

func (d *TDrawGrid) TabStops(Index int32) bool {
    return DrawGrid_GetTabStops(d.instance, Index)
}

func (d *TDrawGrid) SetTabStops(Index int32, value bool) {
    DrawGrid_SetTabStops(d.instance, Index, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (d *TDrawGrid) DockClients(Index int32) *TControl {
    return AsControl(DrawGrid_GetDockClients(d.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (d *TDrawGrid) Controls(Index int32) *TControl {
    return AsControl(DrawGrid_GetControls(d.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (d *TDrawGrid) Components(AIndex int32) *TComponent {
    return AsComponent(DrawGrid_GetComponents(d.instance, AIndex))
}

