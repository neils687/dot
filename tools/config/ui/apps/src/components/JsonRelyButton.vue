<template>
    <div>
        <el-button @click="ShowJsonDialog(objc)">JSON</el-button>
        <el-drawer
                title="JSON textarea!"
                :before-close="handleClose"
                :visible.sync="dialog"
                direction="ltr"
                custom-class="demo-drawer"
                ref="drawer"
        >
            <el-input
                    type="textarea"
                    :autosize="{ minRows: 10, maxRows: 30}"
                    placeholder="请输入内容"
                    v-model="textarea">
            </el-input>
        </el-drawer>
    </div>
</template>

<script lang="ts">
    import Vue from 'vue'
    import {jsonParseRely,makeJsonRely} from "@/components/changeDataStructure/chDS";

    export default Vue.extend({
        name: "JsonRelyButton",
        props: {
            objc: {}
        },
        data() {
            return {
                dialog: false,
                schemaObject: {},
                textarea: ''
            }
        },
        methods: {
            ShowJsonDialog(obj:any){
                this.dialog = true;
                (this as any).objc = obj;
                let data = makeJsonRely(obj);
                this.textarea = JSON.stringify(data,null,4);

            },
            handleClose(done:any){
                try{
                    if(this.textarea){
                        let objct:any = JSON.parse(this.textarea);
                        if(this.inputCheck(objct)){
                            this.$emit('input',jsonParseRely(objct));
                        }else {
                            (this as any).$message.error('json text input error!');
                        }
                    }else{
                        (this as any).$message.error('json text input error!');
                    }
                }catch (e) {
                    (this as any).$message.error('json text input error!');
                }finally {
                    done();
                }
            },
            inputCheck(input:any): Boolean{
                for (let i in input){
                    if (typeof input[i] !== typeof ''){
                        return false;
                    }
                }
                return true;
            }
        }
    })
</script>

<style scoped>

</style>
