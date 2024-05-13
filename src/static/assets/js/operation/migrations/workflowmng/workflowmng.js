$(document).ready(function () {
    order_type = "name"
    //checkbox all
    $("#th_chall").click(function () {
        if ($("#th_chall").prop("checked")) {
            $("input[name=chk]").prop("checked", true);
        } else {
            $("input[name=chk]").prop("checked", false);
        }
    })

    //table 스크롤바 제한
    $(window).on("load resize", function () {
        var vpwidth = $(window).width();
        if (vpwidth > 768 && vpwidth < 1800) {
            $(".dashboard_cont .dataTable").addClass("scrollbar-inner");
            $(".dataTable.scrollbar-inner").scrollbar();
        } else {
            $(".dashboard_cont .dataTable").removeClass("scrollbar-inner");
        }

        setTableHeightForScroll('workflowListTable', 300);
    });
});

function deleteWorkflow() {
    var workflowId = "";
    var count = 0;

    $("input[name='chk']:checked").each(function () {
        count++;
        workflowId = workflowId + $(this).val() + ",";
    });
    workflowId = workflowId.substring(0, workflowId.lastIndexOf(","));

    console.log("workflowId : ", workflowId);
    console.log("count : ", count);

    if (workflowId == '') {
        commonAlert("삭제할 대상을 선택하세요.");
        return false;
    }

    if (count != 1) {
        commonAlert("삭제할 대상을 하나만 선택하세요.");
        return false;
    }

    var url = "/operation/migrations" + "/workflow/del/" + workflowId
    console.log("del workflow url : ", url);

    axios.delete(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        var data = result.data;
        console.log(result);
        console.log(data);
        if (result.status == 200 || result.status == 201) {
            commonAlert(data.message)
            displayWorkflowInfo("DEL_SUCCESS")
        } else {
            commonAlert(result.data.error)
        }
    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

function getWorkflowList(sort_type) {
    console.log(sort_type);
    var url = "/operation/migrations/workflow/list";
    axios.get(url, {
        headers: {
            'Content-Type': "application/json"
        }
    }).then(result => {
        console.log("get Workflow List : ", result.data);
        
        var data = result.data.WorkflowList;

        var html = ""
        var cnt = 0;

        if (data == null) {
            html += '<tr><td class="overlay hidden" data-th="" colspan="5">No Data</td></tr>'

            $("#workflowList").empty()
            $("#workflowList").append(html)

            ModalDetail()
        } else {
            if (data.length) {
                if (sort_type) {
                    cnt++;
                    console.log("check : ", sort_type);
                    data.filter(list => list.Name !== "").sort((a, b) => (a[sort_type] < b[sort_type] ? - 1 : a[sort_type] > b[sort_type] ? 1 : 0)).map((item, index) => (
                        html += addWorkflowRow(item, index)                        
                    ))
                } else {
                    data.filter((list) => list.Name !== "").map((item, index) => (
                        html += addWorkflowRow(item, index)
                    ))
                }

                $("#vpcList").empty()
                $("#vpcList").append(html)

                ModalDetail()
            }
        }

    }).catch((error) => {
        console.warn(error);
        console.log(error.response)
        var errorMessage = error.response.data.error;
        var statusCode = error.response.status;
        commonErrorAlert(statusCode, errorMessage);
    });
}

// VNet목록에 Item 추가
function addWorkflowRow(item, index) {
    console.log("addWorkflowRow " + index);
    console.log(item)
    var html = ""
    html += '<tr onclick="showWorkflowInfo(\'' + item.name + '\');">'
        + '<td class="overlay hidden column-50px" data-th="">'
        + '<input type="hidden" id="sg_info_' + index + '" value="' + item.name + '"/>'
        + '<input type="checkbox" name="chk" value="' + item.name + '" id="raw_' + index + '" title="" /><label for="td_ch1"></label> <span class="ov off"></span>'
        // + '<input type="hidden" value="' + item.systemLabel + '"/>'
        + '</td>'
        + '<td class="btn_mtd ovm" data-th="name">' + item.name + '<span class="ov"></span></td>'
        + '<td class="overlay hidden" data-th="targetModel">' + item.TargetModel + '</td>'
        + '<td class="overlay hidden" data-th="description">' + item.description + '</td>'
        + '</tr>'
    return html;
}

function ModalDetail() {
    $(".dashboard .status_list tbody tr").each(function () {
        var $td_list = $(this),
            $status = $(".server_status"),
            $detail = $(".server_info");
        $td_list.off("click").click(function () {
            $td_list.addClass("on");
            $td_list.siblings().removeClass("on");
            $status.addClass("view");
            $status.siblings().removeClass("on");
            $(".dashboard.register_cont").removeClass("active");
            $td_list.off("click").click(function () {
                if ($(this).hasClass("on")) {
                    console.log("reg ok button click")
                    $td_list.removeClass("on");
                    $status.removeClass("view");
                    $detail.removeClass("active");
                } else {
                    $td_list.addClass("on");
                    $td_list.siblings().removeClass("on");
                    $status.addClass("view");

                    $status.siblings().removeClass("view");
                    $(".dashboard.register_cont").removeClass("active");
                }
            });
        });
    });
}

function displayWorkflowInfo(targetAction) {
    if (targetAction == "REG") {
        // workflow 등록 화면 호출( popup 또는 이동)
        changePage('WorkflowRegForm')
    } else if (targetAction == "REG_SUCCESS") {
        
        getWorkflowList("name");// 등록 성공시 workflow 조회
    } else if (targetAction == "DEL") {
        $('#workflowCreateBox').removeClass("active");
        $('#workflowInfoBox').addClass("view");
        $('#workflowListDiv').removeClass("on");

        var offset = $("#vNetInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 300);

    } else if (targetAction == "DEL_SUCCESS") {
        $('#workflowCreateBox').removeClass("active");
        $('#workflowInfoBox').removeClass("view");
        $('#workflowListDiv').addClass("on");

        var offset = $("#workflowInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);

        getVpcList("name");
    } else if (targetAction == "CLOSE") {
        $('#workflowCreateBox').removeClass("active");
        $('#workflowInfoBox').removeClass("view");
        $('#workflowListDiv').addClass("on");

        var offset = $("#workflowInfoBox").offset();
        $("#TopWrap").animate({ scrollTop: offset.top }, 0);
    }

}

function showWorkflowInfo(workflowId) {
    var url = "/operation/migrations" + "/workflow/" + workflowId;
    console.log("vnet detail URL : ", url)

    return axios.get(url, {

    }).then(result => {
        console.log(result);
        console.log(result.data);
        var data = result.data.WorkflowInfo
        console.log("Show Data : ", data);

        var dtlWorkflowName = data.name;
        var dtlDescription = data.description;
        var dtlTargetModel = data.TargetModel;


        $("#dtlWorkflowName").empty();
        $("#dtlDescription").empty();
        $("#dtlTargetModel").empty();
        
        $("#dtlWorkflowName").val(dtlWorkflowName);
        $("#dtlDescription").val(dtlDescription);
        $("#dtlTargetModel").val(dtlConnectionName);
        
        $('#workflowName').text(dtlWorkflowName)

    }).catch(function (error) {
        console.log("Network detail error : ", error);
    });
}


// function displaySubnetRegModal(isShow) {
//     if (isShow) {
//         $("#subnetRegisterBox").modal();
//         $('.dtbox.scrollbar-inner').scrollbar();
//     } else {
//         $("#vnetCreateBox").toggleClass("active");
//     }
// }



///////////////////////////////////////// Workflow Editor 정의영역 start ///////////////////////////////////
//
// div 설정 : workflowplaceholder
// workflow 저장
//      불러오기 : loadState()
//      저장하기 : saveState()
//      // 
// 설정 : toolbox, controlBar, steps, validator, editor

// workflow event
//      onDefinitionChaged -> refreshValidationStatus(); saveState();
//      changeReadonlyButton : click -> setIsReadonly(), reloadChangeReadonlyButtonText()

//  validation
//      refreshValidationStatus()


//////////////////////////
// mng는 id로 조회해여 보여준다.
// 최초에는 defaultDefinition()

//////////////////////////



// workflowId로 조회
function loadWorkflow(workflowId) {
    
    // alert("get workflow ", workflowId)
    // getCommonCloudConnectionList("mcismng", true)
    console.log("loadState getCommonCloudConnectionList")
	// const state = localStorage[localStorageKey];
	// if (state) {
	// 	return JSON.parse(state);
	// }
	// return {
	// 	definition: getStartDefinition()
	// }
    alert("Workflow loaded")
}
function saveWorkflow() {
    alert("workflow saved")
	// localStorage[localStorageKey] = JSON.stringify({
	// 	definition: designer.getDefinition(),
	// 	undoStack: designer.dumpUndoStack()
	// });
}


function refreshValidationStatus() {
	validationStatusText.innerText = designer.isValid() ? 'Definition is valid' : 'Definition is invalid';
}


const placeholder = document.getElementById('workflowplaceholder');
const localStorageKey = 'sqdMngScreen';
const confState = readOnlyState();
const designer = sequentialWorkflowDesigner.Designer.create(placeholder, confState.definition, confState.configuration);

document.addEventListener('DOMContentLoaded', function() {
    
    console.log("DOMContentLoaded ")
    
    console.log("workflow definition ", confState.definition)

    console.log("workflow configuration ", confState.configuration)
    console.log("placeholder ", placeholder)
    
    designer.onDefinitionChanged.subscribe((newDefinition) => {
        refreshValidationStatus();
        saveWorkflow();
        console.log('the definition has changed', newDefinition);
    });

    console.log("workflow designer created ")

    // workflow 불러오기
    loadWorkflow("w01");
});





///////////////////
// TODO : readonly인데 toolbox가 필요있나?
function toolboxGroup(name) {
	return {
		name,
		steps: [
			createTaskStep(null, 'save', 'Save file'),
			createTaskStep(null, 'text', 'Send email'),
			createTaskStep(null, 'task', 'Create task'),
			createIfStep(null, [], []),
			createContainerStep(null, [])
		]
	};
}

function reloadChangeReadonlyButtonText() {
	changeReadonlyButton.innerText = 'Readonly: ' + (designer.isReadonly() ? 'ON' : 'OFF');
}

