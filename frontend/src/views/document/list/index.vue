<template>
    <div id="document-list">
        <el-table
                :data="tableData"
                style="width: 100%"
                border
        >
            <el-table-column
                    prop="doc_id"
                    label="doc_id"
                    width="100"
            ></el-table-column>

            <el-table-column
                    label="title"
            >
                <template slot-scope="scope">
                    <el-link type="primary" :key="scope.row.doc_id" @click="openDetail(scope.row.doc_id)">{{scope.row.title}}</el-link>
                </template>
            </el-table-column>

            <el-table-column
                    prop="category_name"
                    label="category"
                    width="100"
            ></el-table-column>
            <el-table-column
                    prop="tags"
                    label="tags"
                    width="200"
            >
                <template slot-scope="scope">
                    <el-tag effect="light">tag 1</el-tag>
                    <el-tag effect="light">tag 2</el-tag>
                    <el-tag effect="light">tag 3</el-tag>
                </template>
            </el-table-column>
            <el-table-column
                    prop="author_name"
                    label="author"
                    width="100"
            ></el-table-column>
            <el-table-column
                    prop="post_status"
                    label="status"
                    width="100"
            ></el-table-column>
            <el-table-column
                    prop="create_time"
                    label="time"
                    width="100"
            ></el-table-column>
        </el-table>
        <el-pagination
                background
                layout="total, prev, pager, next"
                :total="1000">
        </el-pagination>
    </div>
</template>

<script>
    export default {
        name: "document_list",
        el: '#document-list',
        data(){
            return {
                tableData: [
                    {
                        doc_id: '1',
                        title: 'test title',
                        category_name: 'api',
                        author_name: 'author name',
                        post_status: '1',
                        create_time: '2019-01-01'
                    },
                    {
                        doc_id: '2',
                        title: 'test title',
                        category_name: 'api',
                        author_name: 'author name',
                        post_status: '1',
                        create_time: '2019-01-01'
                    }
                ]
            }
        },
        methods: {
            openDetail: function (docId) {
                this.$emit('custom-tab-open', {
                    tabType: 'detail',
                    tabData: {
                        docId: docId,
                    }
                })
            }
        }
    }
</script>

<style scoped>
    .el-table {
        margin-bottom: 15px;
    }
</style>