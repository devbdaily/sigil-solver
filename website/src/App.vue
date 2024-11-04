<script>
import PieceButton from "./components/PieceButton.vue";
import Block from "./components/icons/Block.vue";
import IPiece from "./components/icons/IPiece.vue";
import JPiece from "./components/icons/JPiece.vue";
import LPiece from "./components/icons/LPiece.vue";
import OPiece from "./components/icons/OPiece.vue";
import SPiece from "./components/icons/SPiece.vue";
import TPiece from "./components/icons/TPiece.vue";
import ZPiece from "./components/icons/ZPiece.vue";
import PieceDefs from "./pieces";
import axios from "axios";

export default {
    data() {
        return {
            input: {
                board: {
                    width: 4,
                    height: 3,
                },
                pieces: ["J", "J", "Z"],
            },
            result: null,
            error: null,
            colors: [
                "fill-red-500",
                "fill-orange-500",
                "fill-amber-500",
                "fill-yellow-500",
                "fill-lime-500",
                "fill-green-500",
                "fill-emerald-500",
                "fill-teal-500",
                "fill-cyan-500",
                "fill-sky-500",
                "fill-blue-500",
                "fill-indigo-500",
                "fill-violet-500",
                "fill-purple-500",
                "fill-fuchsia-500",
                "fill-pink-500",
                "fill-rose-500",
            ],
        }
    },
    methods: {
        solvePuzzle() {
            var body = {
                board: this.input.board,
                pieces: [],
            }

            for (var piece of this.input.pieces) {
                body.pieces.push(PieceDefs[piece])
            }

            axios.post("/solve", body)
            .then((response) => {
                this.result = response.data
            })
            .catch((error) => {
                this.error = error.response.data.error
            })
        },
        clearData() {
            this.input.board.width = 1
            this.input.board.height = 1
            this.input.pieces = []
            this.result = null
            this.error = null
        },
        addPiece(shape) {
            this.input.pieces.push(shape)
        },
        getPieceComponent(piece) {
            switch (piece) {
                case "I":
                    return IPiece
                case "J":
                    return JPiece
                case "L":
                    return LPiece
                case "O":
                    return OPiece
                case "S":
                    return SPiece
                case "T":
                    return TPiece
                case "Z":
                    return ZPiece
            }
        },
    },
    components: {
        Block,
        PieceButton,
        IPiece,
        JPiece,
        LPiece,
        OPiece,
        SPiece,
        TPiece,
        ZPiece,
    },
}
</script>

<template>
    <main class="container mx-auto space-y-2">
        <div class="w-32">
            <h2 class="text-2xl font-bold">Board Info</h2>
            <div class="space-x-2 flex">
                <label for="width" class="font-bold flex-1">Width:</label>
                <input type="number" name="width" id="width" min="1" max="16" v-model="input.board.width" class="border flex-none">
            </div>
            <div class="space-x-2 flex">
                <label for="height" class="font-bold flex-1">Height:</label>
                <input type="number" name="height" id="height" min="1" max="16" v-model="input.board.height" class="border flex-none">
            </div>
        </div>
        <div>
            <h2 class="text-2xl font-bold">Piece Info</h2>
            <div class="space-x-2">
                <PieceButton @click="addPiece('I')">
                    <IPiece class="fill-red-500" />
                </PieceButton>
                <PieceButton @click="addPiece('J')">
                    <JPiece class="fill-red-500" />
                </PieceButton>
                <PieceButton @click="addPiece('L')">
                    <LPiece class="fill-red-500" />
                </PieceButton>
                <PieceButton @click="addPiece('O')">
                    <OPiece class="fill-red-500" />
                </PieceButton>
                <PieceButton @click="addPiece('S')">
                    <SPiece class="fill-red-500" />
                </PieceButton>
                <PieceButton @click="addPiece('T')">
                    <TPiece class="fill-red-500" />
                </PieceButton>
                <PieceButton @click="addPiece('Z')">
                    <ZPiece class="fill-red-500" />
                </PieceButton>
            </div>
            <div>
                <h3 class="text-lg font-bold">Selected Pieces</h3>
                <div class="space-x-2">
                    <component v-for="(piece, index) in input.pieces" :key="index" :is="getPieceComponent(piece)" class="inline-block" :class="[colors[index]]"></component>
                </div>
            </div>
        </div>
        <div class="space-x-2">
            <button @click="clearData" class="border rounded bg-slate-200 px-2 py-1">
                Clear
            </button>
            <button @click="solvePuzzle" class="border rounded bg-blue-500 text-white px-2 py-1">
                Submit
            </button>
        </div>
        <div>
            <h2 class="text-2xl font-bold">Result</h2>
            <div v-if="error !== null" v-text="error" class="text-red-500"></div>
            <div class="relative" v-if="result !== null">
                <template v-for="(piece, pIndex) in result.pieces">
                    <Block v-for="(block, bIndex) in piece.blocks" :key="bIndex" :width="input.board.width" :height="input.board.height" :x="block.x" :y="block.y" :class="[colors[pIndex]]" />
                </template>
            </div>
        </div>
    </main>
</template>
