<template>
  <div class="bezier-editor-container">
    <div class="controls">
      <button class="bg-neutral-800 p-2 border rounded-lg" @click="addCurveSegment">
        + Segment
      </button>
    </div>
    <div ref="konvaContainer" class="border border-neutral-950 bg-neutral-800 w-[1002px]"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue'
import Konva from 'konva'
import type { Layer } from 'konva/lib/Layer'
import { useVModel } from '@vueuse/core'
const props = defineProps<{
  modelValue: BezierSegment[]
}>()
const emit = defineEmits(['update:modelValue'])

// --- TYPE DEFINITIONS ---
export interface Point {
  x: number
  y: number
}

export interface BezierSegment {
  p0: Point
  p1: Point // Control Point 1
  p2: Point // Control Point 2
  p3: Point
}

// --- COMPONENT STATE (REFS) ---
const konvaContainer = ref<HTMLDivElement | null>(null)
const stage = ref<Konva.Stage | null>(null)
const layer = ref<Konva.Layer | null>(null)
const gridLayer = ref<Konva.Layer | null>(null)
const gridSize = 10

// Restructured data model: An array of 4-point bezier curve objects.
const curves = useVModel(props, 'modelValue', emit)
// const curves = ref<BezierSegment[]>([
//   {
//     p0: { x: 0, y: 50 },
//     p1: { x: 20, y: 50 },
//     p2: { x: 40, y: 50 },
//     p3: { x: 60, y: 50 },
//   },
// ])

// --- LIFECYCLE HOOK ---
onMounted(() => {
  // Wait for the container to be fully rendered to get correct dimensions
  nextTick(() => {
    initializeKonva()
    drawScene()
  })
})

// --- KONVA INITIALIZATION ---
const initializeKonva = () => {
  if (!konvaContainer.value) return

  const width = konvaContainer.value.clientWidth || 600
  const height = 100 // Fixed height

  stage.value = new Konva.Stage({
    container: konvaContainer.value,
    width: width,
    height: height,
  })

  layer.value = new Konva.Layer()
  gridLayer.value = new Konva.Layer()
  stage.value.add(gridLayer.value as Layer)
  stage.value.add(layer.value as Layer)

  // Handle window resizing
  window.addEventListener('resize', () => {
    if (stage.value && konvaContainer.value) {
      const newWidth = konvaContainer.value.clientWidth
      stage.value.width(newWidth)
    }
  })
}

// --- DRAWING LOGIC ---
/**
 * Clears and redraws the entire Konva scene from the `curves` data.
 * This should only be called on initialization or when the number of curves changes.
 */
const drawScene = () => {
  if (!layer.value) return
  layer.value.destroyChildren() // Clear previous drawings

  // Draw each curve segment
  curves.value.forEach((curve) => {
    // Draw the main Bézier Curve.
    // The sceneFunc reads from the reactive `curves` ref.
    // This means calling `layer.batchDraw()` is enough to update its shape.
    const bezierCurve = new Konva.Shape({
      stroke: 'orange',
      strokeWidth: 2,
      sceneFunc: (context, shape) => {
        context.beginPath()
        context.moveTo(curve.p0.x, curve.p0.y)
        context.bezierCurveTo(
          curve.p1.x,
          curve.p1.y,
          curve.p2.x,
          curve.p2.y,
          curve.p3.x,
          curve.p3.y,
        )
        context.fillStrokeShape(shape)
      },
    })
    layer.value?.add(bezierCurve)

    const guideLine1 = new Konva.Shape({
      stroke: 'grey',
      strokeWidth: 1,
      dash: [4, 4],
      sceneFunc: (context, shape) => {
        context.beginPath()
        context.moveTo(curve.p0.x, curve.p0.y)
        context.lineTo(curve.p1.x, curve.p1.y)
        context.fillStrokeShape(shape)
      },
    })
    const guideLine2 = new Konva.Shape({
      stroke: 'grey',
      strokeWidth: 1,
      dash: [4, 4],
      sceneFunc: (context, shape) => {
        context.beginPath()
        context.moveTo(curve.p3.x, curve.p3.y)
        context.lineTo(curve.p2.x, curve.p2.y)
        context.fillStrokeShape(shape)
      },
    })
    layer.value?.add(guideLine1, guideLine2)

    const width = stage.value?.width() ?? 0
    const height = stage.value?.height() ?? 0

    gridLayer.value?.destroyChildren()
    for (let i = 0; i < width / gridSize; i++) {
      gridLayer.value?.add(
        new Konva.Line({
          points: [Math.round(i * gridSize) + 0.5, 0, Math.round(i * gridSize) + 0.5, height],
          stroke: '#ffffff55',
          strokeWidth: 0.5,
        }),
      )
    }

    for (let j = 0; j < height / gridSize; j++) {
      gridLayer.value?.add(
        new Konva.Line({
          points: [0, Math.round(j * gridSize), width, Math.round(j * gridSize)],
          stroke: '#ffffff55',
          strokeWidth: 0.5,
        }),
      )
    }
  })

  // Draw all the interactive handles on top so they are not obscured
  curves.value.forEach((curve, curveIndex) => {
    // Note: We avoid creating duplicate handles for connected anchor points.
    // The 'start' handle of a curve is only drawn if it's the very first curve.
    if (curveIndex === 0) {
      createHandle(curve.p0)
    }
    createHandle(curve.p1)
    createHandle(curve.p2)
    createHandle(curve.p3)
  })

  layer.value.batchDraw()
}

/**
 * Creates a single draggable handle.
 * @param point The point object {x, y} to create a handle for.
 */
const createHandle = (point: Point) => {
  if (!layer.value || !stage.value) return

  // Determine if the point is a main anchor or a control point by checking its reference
  const isAnchor = curves.value.some((c) => c.p0 === point || c.p3 === point)

  const handle = new Konva.Circle({
    x: point.x,
    y: point.y,
    radius: isAnchor ? 7 : 5,
    fill: isAnchor ? 'white' : 'gray',
    stroke: 'black',
    strokeWidth: 1,
    draggable: true,
    hitStrokeWidth: 10, // Easier to grab
    dragBoundFunc: function (pos) {
      const stageWidth = stage.value!.width()
      const stageHeight = stage.value!.height()

      // Get the current scale and position of the stage
      const scale = stage.value!.scaleX() // assuming uniform scaling
      const stagePos = stage.value!.position()

      // Transform stage boundaries into the layer's coordinate system
      const minX = -stagePos.x / scale
      const maxX = (stageWidth - stagePos.x) / scale
      const minY = -stagePos.y / scale
      const maxY = (stageHeight - stagePos.y) / scale

      // Clamp the handle's position within these calculated boundaries
      const newX = Math.max(minX, Math.min(pos.x, maxX))
      const newY = Math.max(minY, Math.min(pos.y, maxY))

      return { x: newX, y: newY }
    },
  })

  handle.on('dragmove', (e) => {
    // 1. Directly update the coordinates of the underlying data point.
    point.x = e.target.x()
    point.y = e.target.y()

    const newX = Math.round(point.x / gridSize) * gridSize
    const newY = Math.round(point.y / gridSize) * gridSize

    point.x = newX
    point.y = newY
    handle.x(newX)
    handle.y(newY)

    // 2. Redraw the layer. This is efficient. Konva re-runs the sceneFuncs
    // for the curves, which read the new data, and updates the lines.
    // This happens without destroying and recreating the objects, fixing the bug.
    layer.value?.batchDraw()
  })

  // Basic cursor changes for better UX
  handle.on('mouseenter', () => (stage.value!.container().style.cursor = 'grab'))
  handle.on('mouseleave', () => (stage.value!.container().style.cursor = 'default'))
  handle.on('mousedown', () => (stage.value!.container().style.cursor = 'grabbing'))
  handle.on('mouseup', () => (stage.value!.container().style.cursor = 'grab'))

  layer.value.add(handle)
}

// --- USER ACTIONS ---
/**
 * Adds a new curve segment connected to the last one.
 */
const addCurveSegment = () => {
  if (!curves.value) {
    curves.value = [
      {
        p0: {
          x: 0,
          y: 50,
        },
        p1: {
          x: 20,
          y: 50,
        },
        p2: {
          x: 40,
          y: 50,
        },
        p3: {
          x: 60,
          y: 50,
        },
      },
    ]
    return
  }

  const lastCurve = curves.value[curves.value.length - 1]
  if (!lastCurve) {
    return
  }

  // The new curve's start point IS the last curve's end point.
  // This linking by reference is the key to keeping them connected.
  const newStartPoint = lastCurve?.p3

  const newSegment: BezierSegment = {
    p0: newStartPoint,
    p1: { x: newStartPoint.x + 20, y: newStartPoint.y },
    p2: { x: newStartPoint.x + 40, y: newStartPoint.y },
    p3: { x: newStartPoint.x + 60, y: newStartPoint.y },
  }

  curves.value.push(newSegment)

  // Since the watcher is gone, we must manually trigger a full redraw
  // to create the new shapes for the new segment.
  drawScene()
}
</script>

<style scoped></style>
