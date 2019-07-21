<template>
    <div class="tab-content">
        <el-table
                :data="tableData"
                style="width: 100%"
                border
        >
            <el-table-column
                    prop="doc_id"
                    label="id"
                    width="50"
            ></el-table-column>

            <el-table-column
                    label="title"
            >
                <template slot-scope="scope">
                    <el-link type="primary" :key="scope.row.doc_id" @click="openDetail(scope.row)">{{scope.row.title}}</el-link>
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
    import apis from '@/api/apis'

    export default {
        name: "document_list",
        data(){
            return {
                tableData: [
                    {
                        doc_id: '1',
                        title: 'test title',
                        category_name: 'api',
                        author_name: 'author name',
                        spec_url: '',
                        post_status: '1',
                        create_time: '2019-01-01'
                    },
                    {
                        doc_id: '2',
                        title: 'test title',
                        category_name: 'api',
                        author_name: 'author name',
                        spec_url: '',
                        post_status: '1',
                        create_time: '2019-01-01'
                    }
                ]
            }
        },
        methods: {
            openDetail: function (row) {
                this.$emit('custom-tab-open', {
                    tabType: 'detail',
                    tabData: {
                        doc_id: row.doc_id,
                        title: row.title,
                        spec_url: row.spec_url
                    }
                })
            }
        },
        mounted() {
            apis.docList(this, 1, 20, function (data, err) {
                console.log(data);
                console.log(err);
            })
        }
    }
</script>

<style scoped>
    .el-table {
        margin-bottom: 15px;
    }

    .el-tag {
        margin-right: 5px;
    }

</style>