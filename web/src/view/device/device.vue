/* eslint-disable no-debugger */
<template>
    <div>
        <el-form :model="form3" label-width="80px" ref="form">

            <el-table
                    :data="tableData"
                    border
                    ref="multipleTable"
                    stripe
                    style="width: 100%"
                    tooltip-effect="dark"
            >
                <el-table-column label="设备名称" prop="DeviceName" v-model="form3.DeviceName" width="120"></el-table-column>
                <el-table-column label="设备类型" prop="DeviceType" v-model="form3.DeviceType" width="120"></el-table-column>
                <el-table-column label="节点名称" prop="NodeName" v-model="form3.NodeName"  width="120"></el-table-column>
                <!--<el-table-column label="设备描述" prop="DeviceSpec" width="120"></el-table-column>
                <el-table-column label="设备状态" prop="Status" width="120"></el-table-column>-->
                <el-table-column label="创建时间" prop="CreateTime" width="120"></el-table-column>
                <el-table-column label="按钮组" min-width="160">
                    <template slot-scope="scope">
                        <el-button @click="getDeviceFile(scope.row)" size="small" type="text">获取设备文件</el-button>
                    </template>
                </el-table-column>
              <!--  <el-table-column label="按钮组" min-width="160">
                    <template slot-scope="scope">
                        &lt;!&ndash;<el-button @click="getDeviceFile()" size="big" type="text">test</el-button>&ndash;&gt;
                        <el-button @click="getDeviceFile(scope.row)" size="small" type="text">变更</el-button>
                        <el-popover placement="top" width="160" v-model="scope.row.visible">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin: 0">
                                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">确定</el-button>
                            </div>
                            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                        </el-popover>
                    </template>
                </el-table-column>-->
            </el-table>
        </el-form>
        <el-form :inline="true" :model="form" label-width="80px" v-show="showRentPrise">
            <el-row>
                <el-col :span="3"><label for="">红灯</label></el-col>
                <el-col :span="10">
                    <el-switch v-model="form.red"></el-switch>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3"><label for="">黄灯</label></el-col>
                <el-col :span="10">
                    <el-switch v-model="form.yellow">
                    </el-switch>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3"><label for="">绿灯</label></el-col>
                <el-col :span="10">
                    <el-switch v-model="form.green"></el-switch>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3"><label for="">设备文件</label></el-col>
                <el-col :span="10">
                    <el-input type="textarea" v-model="form.desc"></el-input>
                </el-col>
            </el-row>

            <el-row type="flex" justify="center">
                <el-col :span="13">
                    <el-button @click="onSubmit" type="primary">提交</el-button>
                    <el-button>取消</el-button>
                </el-col>

            </el-row>
        </el-form>
        <el-form :inline="true" :model="form1" label-width="80px" v-if="showPrise">
            <el-row>
                <el-col :span="3"><label for="">温度:</label></el-col>
                <el-col :span="1">
                    <!-- <el-switch v-model="form.temp"></el-switch>-->
                    <el-input v-model="form1.temp" placeholder=""></el-input>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3"><label for="">设备文件</label></el-col>
                <el-col :span="10">
                    <el-input type="textarea" v-model="form1.desc"></el-input>
                </el-col>
            </el-row>
            <el-row type="flex" justify="center">
                <el-col :span="13">
                    <el-button @click="onSubmit" type="primary">提交</el-button>
                    <el-button>取消</el-button>
                </el-col>

            </el-row>
        </el-form>


    </div>
</template>

<script>
    import {
        getDeviceFile,
        updateDevice,
        getDeviceDetails
    } from "@/api/customer";
    import infoList from "@/components/mixins/infoList";

    export default {
        name: "Node",
        mixins: [infoList],
        data() {
            return {
                listApi: getDeviceDetails,
                dialogFormVisible: false,
                visible: false,
                type: "",
                form1: {
                    temp: "",
                    desc: ""
                },
                form: {
                    red: "",
                    green: "",
                    yellow: "",
                    desc: ""
                },
                form3: {
                    DeviceName: "",
                    DeviceType: "",
                    NodeName: "",
                    
                },
                showRentPrise: false,
                showPrise: false
            };
        },
        methods: {
            // eslint-disable-next-line no-,no-unused-vars
            async getDeviceFile(row) {
                // eslint-disable-next-line no-debugger
                debugger
                const res = await getDeviceFile({
                    DeviceName: row.DeviceName,
                    DeviceType: row.DeviceType,
                    NodeName: row.NodeName
                });
                // eslint-disable-next-line no-debugger
                if (row.DeviceType == "temperature-model") {
                    this.showPrise = true
                }
                if (row.DeviceType == "traffic-light") {
                    this.showRentPrise = true
                }
                var desc = res.data
                if (res.code == 0 && row.DeviceType == "traffic-light") {
                    // eslint-disable-next-line no-debugger
                    debugger

                    // eslint-disable-next-line no-

                    var obj = res.data;
                    var colors = obj.status.twins
                    var color_red = colors[0]
                    var color_green = colors[1]
                    var color_yellow = colors[2]
                    if (color_red.reported.value == "ON") {
                        this.form.red = true
                    } else {
                        this.form.red = false
                    }
                    if (color_green.reported.value == "ON") {
                        this.form.green = true
                    } else {
                        this.form.green = false
                    }
                    if (color_yellow.reported.value == "ON") {
                        this.form.yellow = true
                    } else {
                        this.form.yellow = false
                    }
                    // eslint-disable-next-line no-debugger
                    debugger

                    var strOfDevice = JSON.stringify(color_red.desired.value + color_yellow.desired.value + color_green.desired.value)
                    this.form.desc = strOfDevice + JSON.stringify(desc);
                    // var red= ""
                    // var green = ""
                    // var yellow = ""
                } else if (res.code == 0 && row.DeviceType == "temperature-model") {
                    // eslint-disable-next-line no-redeclare,no-debugger
                    debugger
                    // eslint-disable-next-line no-redeclare
                    var obj = res.data.status.twins;
                    this.form1.temp = obj[0].reported.value
                    // eslint-disable-next-line no-redeclare
                    var strOfDevice = JSON.stringify(obj[0].desired.value)
                    this.form1.desc=strOfDevice+JSON.stringify(desc)
                }
            },
            async onSubmit() {

                var redValue
                var greenValue
                var yellowValue
                // eslint-disable-next-line no-undef
                var deiveName=this.form3.DeviceName
                // eslint-disable-next-line no-undef
                var deviceType=this.form3.DeviceType
                // eslint-disable-next-line no-undef
                var nodeName=this.form3.NodeName
                if (this.form.red == true) {
                    redValue = "ON"
                } else {
                    redValue = "OFF"
                }
                if (this.form.green == true) {
                    greenValue = "ON"
                } else {
                    greenValue = "OFF"
                }
                if (this.form.yellow == true) {
                    yellowValue = "ON"
                } else {
                    yellowValue = "OFF"
                }

                var redobj = {
                    "name": "red",
                    "value": redValue
                }
                // eslint-disable-next-line no-unused-vars
                var greenobj = {
                    "name": "green",
                    "value": greenValue
                }
                // eslint-disable-next-line no-unused-vars
                var yelloobj = {
                    "name": "yellow",
                    "value": yellowValue
                }
                var Data = {
                    "Data": [redobj,
                        greenobj,
                        yelloobj],
                    "deiveName":deiveName,
                    "deviceType":deviceType,
                    "nodeName":nodeName

                }

                const res = await updateDevice(Data);

                if (res.code == 0) {
                    this.dialogFormVisible = true;
                }

            },
        },
        created() {
            this.getTableData();
        }
    };
</script>

<style>
</style>