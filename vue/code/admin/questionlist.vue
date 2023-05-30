<script>
    export default {
        components: {
            'QuestionEditor': () => import("/vue/questioneditor.js"),
        },
        data() {
            return {
                title: 'Questions',
                breadcrumb: [{
                    text: 'Question',
                },
                {
                    text: 'All',
                    active: true,
                },
                ],
                columns: [
                    { key: "ID", sortable: true, sortkey: "id" },
                    { key: "Question", sortable: true, sortkey: "question" },
                ],
                table: false,
                error: "",
                message: "",
                actions: [
                    // {
                    //     key: "import",
                    //     icon: "ri-upload-line",
                    //     text: "Import"
                    // }
                ],
                new: {
                    ID: 0,
                    Question: "",
                },
                showsubquestionsof: false
            }
        },
        methods: {
            onAction(action, arg) {
                switch (action) {
                    case 'loading':
                        this.loading = true
                        this.table = arg
                        break;
                    case 'add_new':
                        this.table.editing_item = Object.assign({}, this.new)
                        break;
                    default:
                        break;
                }
            },
            onActionDone(data) {
                this.loading = false
            },
            showsubquestions(record) {
                if(record && record.ID > 0){
                    this.showsubquestionsof = record
                }
            }
        },
        template: `{{{template}}}`
    }
</script>

<template>
        <div class="col-12">
            <div class="card">
                <List api="questions" :columns="columns" title_column="Code" :can_select="true" :can_export="true"
                    :can_import="true" :actions="actions" @done="onActionDone" @onaction="onAction"> <!--fullscreen-->
                    <template v-slot:SubQuestionCount="{ row, col }">
                        <div class="form-check form-switch">
                            <button class="btn btn-primary" type="button" @click="showsubquestions(row)">
                                {{row.SubQuestionCount}}
                            </button>
                        </div>
                    </template>
                    <template v-slot:editor="editing_item">
                        <QuestionEditor v-if="editing_item.item" :value="editing_item.item" @input="editing_item.submit">
                        </QuestionEditor>
                    </template>
                </List>
            </div>
        </div>
    </div>
</template>