<template>
    <div class="tab-content">
        <div style="margin-bottom: 15px;">
            <el-input placeholder="请输入搜索内容" v-model="search" class="input-with-select" @keyup.enter.native="handleSearch">
                <el-button slot="append" icon="el-icon-search" @click="handleSearch"></el-button>
            </el-input>
        </div>

        <el-table
                :data="items"
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

<!--            tags-->
<!--            <el-table-column-->
<!--                    prop="tags"-->
<!--                    label="tags"-->
<!--                    width="200"-->
<!--            >-->
<!--                <template slot-scope="scope">-->
<!--                    <el-tag effect="light">tag 1</el-tag>-->
<!--                    <el-tag effect="light">tag 2</el-tag>-->
<!--                    <el-tag effect="light">tag 3</el-tag>-->
<!--                </template>-->
<!--            </el-table-column>-->

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
                    width="200"
            ></el-table-column>
        </el-table>
        <el-pagination
                background
                layout="total, prev, pager, next"
                hide-on-single-page
                @current-change="handleCurrentChange"
                :page-size="pageSize"
                :total="count">
        </el-pagination>
    </div>
</template>

<script>
    import apis from '@/api/apis'
    import moment from "moment"

    export default {
        name: "document_list",
        data(){
            return {
                pageSize: 20,
                count: 0,
                search: '',
                items: []
            }
        },
        methods: {
            openDetail: function (row) {
                this.$emit('custom-tab-open', {
                    tabType: 'detail',
                    tabData: {
                        doc_id: row.doc_id,
                        title: row.title,
                        spec_url: apis.docDetailSpecUrl(row.doc_id)
                    }
                })
            },
            handleSearch(){
                this.handleCurrentChange(1)
            },
            handleCurrentChange: function (page) {
                this.items = [];
                apis.docList(this, page, this.pageSize, this.search.trim(), (data, err) => {
                    if (err) {
                        return
                    }

                    this.count = data.data.count;

                    let items = data.data.items;

                    for (let i in items) {
                        let item = items[i];
                        let unixCreateTime = item.create_time;
                        let mTime = moment.unix(unixCreateTime);
                        let strTime  =mTime.format("YYYY-MM-DD HH:mm:ss");

                        let postStatus = "unknown";
                        const PostStatusNotPublished = 0;
                        const PostStatusPublished = 1;
                        const PostStatusDeleted = 2;
                        switch (item.post_status) {
                            case PostStatusNotPublished:
                                postStatus = "not_published";
                                break;
                            case PostStatusPublished:
                                postStatus = "published";
                                break;

                            case PostStatusDeleted:
                                postStatus = "deleted";
                                break;

                            default:
                                break

                        }

                        this.items.push({
                            doc_id: item.doc_id,
                            title: item.title,
                            category_name: item.category_name,
                            author_name: item.author_name,
                            spec_url: item.spec_url,
                            post_status: postStatus,
                            create_time: strTime
                        });
                    }


                    // eslint-disable-next-line no-console
                    console.log(data, err, this.items);
                });
            }
        },
        mounted() {
            this.handleCurrentChange(1)
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

    .tab-content .input-with-select .el-input-group__prepend {
        background-color: #fff !important;
    }

</style>