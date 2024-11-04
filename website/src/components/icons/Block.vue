<script>
export default {
    data() {
        return {
            colors: [
                "fill-red-500",
                "fill-orange-500",
                "fill-yellow-500",
                "fill-lime-500",
                "fill-green-500",
                "fill-teal-500",
                "fill-cyan-500",
                "fill-blue-500",
                "fill-purple-500",
                "fill-fuchsia-500",
                "fill-pink-500",
            ],
        }
    },
    props: {
        width: Number,
        height: Number,
        x: Number,
        y: Number,
    },
}
</script>

<template>
    <svg :width="width*25" :height="height*25" class="absolute">
        <rect :x="x*25" :y="25*(height-(y+1))" width="25" height="25"></rect>
    </svg>
</template>
