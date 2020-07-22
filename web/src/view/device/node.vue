<template>
    <div>
        <el-form :model="form" label-width="80px" ref="form">

            <el-table
                    :data="tableData"
                    border
                    ref="multipleTable"
                    stripe
                    style="width: 100%"
                    tooltip-effect="dark"
            >
                <el-table-column label="节点名称" prop="NodeName" width="120"></el-table-column>
                <!--<el-table-column label="设备类型" prop="customerPhoneData" width="120"></el-table-column>
                <el-table-column label="设备名称" prop="sysUserId" width="120"></el-table-column>
                <el-table-column label="设备状态" prop="sysUserId" width="120"></el-table-column>-->
            </el-table>



        </el-form>


    </div>
</template>

<script>
    import {
        getNodes
    } from "@/api/customer";
    import infoList from "@/components/mixins/infoList";

    export default {
        name: "Node",
        mixins: [infoList],
        data() {
            return {
                listApi: getNodes,
                dialogFormVisible: false,
                visible: false,
                type: "",
                form1: {
                    customerName: "",
                    customerPhoneData: ""
                },
                form: {
                    red: "",
                    green: "",
                    yellow: "",
                    desc: ""
                }
            };
        },

        created() {
            this.getTableData();
        }
    };
</script>

<style>
</style>