<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item>
                    <el-button @click="openDialog" type="primary">新增定时任务</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-table
                :data="tableData"
                border
                ref="multipleTable"
                stripe
                style="width: 100%"
                tooltip-effect="dark"
        >
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column label="名称" prop="name" width="120"></el-table-column>
            <el-table-column label="cron表达式" prop="cronexpression" width="120"></el-table-column>
            <el-table-column label="方法名" prop="funcname" width="120"></el-table-column>
            <el-table-column label="备注" prop="note" width="120"></el-table-column>
            <el-table-column label="过期时间" width="180">
                <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
            </el-table-column>
            <el-table-column label="按钮组" min-width="160">
                <template slot-scope="scope">
                    <el-button @click="getDeviceInfo()" size="big" type="text">修改</el-button>
                    <el-button @click="updateCustomer(scope.row)" size="small" type="text">保存</el-button>
                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                            <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">确定</el-button>
                        </div>
                        <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                    </el-popover>
                </template>
            </el-table-column>
        </el-table>

      <!--  <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :style="{float:'right',padding:'20px'}"
                :total="total"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
                layout="total, sizes, prev, pager, next, jumper"
        ></el-pagination>-->



        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="客户">
            <el-form :inline="true" :model="form1" label-width="80px">
                <el-form-item label="客户名">
                    <el-input autocomplete="off" v-model="form1.customerName"></el-input>
                </el-form-item>
                <el-form-item label="客户电话">
                    <el-input autocomplete="off" v-model="form1.customerPhoneData"></el-input>
                </el-form-item>




            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>
        <!--<div class="tips"> 在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示</div>-->




      <!--  <el-form :model="form" label-width="80px" ref="form">

            <el-row>
                <el-col :span="3"><label for="">红灯</label></el-col>
                <el-col :span="10">
                    <el-switch v-model="form.red"></el-switch>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3"><label for="">绿灯</label></el-col>
                <el-col :span="10">
                    <el-switch v-model="form.green"></el-switch>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3"><label for="">黄灯</label></el-col>
                <el-col :span="10">
                    <el-switch v-model="form.yellow"></el-switch>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="3" ><label for="">device文件</label></el-col>
                <el-col :span="10"  >
                    <el-input type="textarea" v-model="form.desc"></el-input>
                </el-col>
            </el-row>



            <el-row type="flex" justify="center">
                <el-col :span="13">
                    <el-button @click="getDeviceInfo()" size="small" type="text">ccccc</el-button>
                    &lt;!&ndash;                      <el-button @click="onSubmit" type="primary">立即创建</el-button>&ndash;&gt;
                    <el-button>取消</el-button>
                </el-col>

            </el-row>

        </el-form>-->




    </div>
</template>

<script>
    import {
        getCronLists
    } from "@/api/customer";
    import { formatTimeToStr } from "@/utils/data";
    import infoList from "@/components/mixins/infoList";

    export default {
        name: "Customer",
        mixins: [infoList],
        data() {
            return {
                listApi: getCronLists,
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
        filters: {
            formatDate: function(time) {
                if (time != null && time != "") {
                    var date = new Date(time);
                    return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
                } else {
                    return "";
                }
            }
        },
       /* methods: {
            async getDeviceInfo() {
                const res = await getDeviceFile();
                this.type = "update";
                if (res.code == 0) {
                    var desc = res.data[0]


                    var obj = res.data[0];
                    var colors = obj.status.twins
                    var color_red = colors[0]
                    this.form.red=colors[0]
                    var color_green = colors[1]
                    this.form.green=colors[1]
                    var color_yellow = colors[2]
                    this.form.yellow=colors[2]

                    var strOfDevice = JSON.stringify(color_red.desired.value + color_green.desired.value + color_yellow.desired.value )
                    this.form.desc = strOfDevice  + JSON.stringify(desc);

                    // var red= ""
                    // var green = ""
                    // var yellow = ""

                }
            },
            async updateCustomer(row) {
                // eslint-disable-next-line no-debugger
                debugger
                const res = await getExaCustomer({ ID: row.ID });
                this.type = "update";
                if (res.code == 0) {
                    // eslint-disable-next-line no-debugger
                    debugger
                    this.form1 = res.data.customer;
                    this.dialogFormVisible = true;
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.form1 = {
                    customerName: "",
                    customerPhoneData: ""
                };
            },
            async deleteCustomer(row) {
                this.visible = false;
                const res = await deleteExaCustomer({ ID: row.ID });
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "删除成功"
                    });
                    this.getTableData();
                }
            },
            async enterDialog() {
                let res;
                switch (this.type) {
                    case "create":
                        res = await createExaCustomer(this.form1);
                        break;
                    case "update":
                        res = await updateExaCustomer(this.form1);
                        break;
                    default:
                        res = await createExaCustomer(this.form1);
                        break;
                }

                if (res.code == 0) {
                    this.closeDialog();
                    this.getTableData();
                }
            },
            openDialog() {
                this.type = "create";
                this.dialogFormVisible = true;
            }
        },*/
        created() {
            this.getTableData();
        }
    };
</script>

<style>
</style>