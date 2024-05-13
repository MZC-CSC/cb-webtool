// Task Component 전용 js
// createXxxStep
// appendXxxBox
// step 안에서  appendTitlebox, appendPath

function createTaskStep(id, type, name, properties) {
	return {
		id,
		componentType: 'task',
		type,
		name,
		properties: properties || {}
	};
}

function createIfStep(id, _true, _false) {
	return {
		id,
		componentType: 'switch',
		type: 'if',
		name: 'If',
		branches: {
			'true': _true,
			'false': _false
		},
		properties: {}
	};
}

function createContainerStep(id, steps) {
	return {
		id,
		componentType: 'container',
		type: 'loop',
		name: 'Loop',
		properties: {},
		sequence: steps
	};
}

function appendTaskPropertiesbox(root, step, isReadonly, editorContext){

	const item = document.createElement('div');

	// console.log("appendTaskPropertiesbox")
	// console.log(step)

	// // 기본 상항으로 사용할 provider와 option

	// // task 별 edit 화면이 정의되어 있음.
	// // 

	// // provider : input
	// const providerInput = document.createElement('input');
	// providerInput.type = 'text';
	// providerInput.name = 'inputParam';
	// providerInput.value = step.properties.provider;
	// if (isReadonly) {
	// 	providerInput.setAttribute('disabled', 'disabled');
	// }

	// const providerLabel = document.createElement('label');
	// providerLabel.textContent = "Provider";
	// providerLabel.htmlFor = providerInput.name;
		
	// item.append(providerLabel);
	// item.append(providerInput);

	// // options : text area => json format
	// const options = document.createElement('textarea');
	// options.name = 'optionParams';
	// options.rows = 4;
	// options.cols = 10
	// options.value = step.properties.options;
	// if (isReadonly) {
	// 	options.setAttribute('disabled', 'disabled');
	// }

	// const optionsLabel = document.createElement('label');
	// optionsLabel.textContent = "Options";
	// optionsLabel.htmlFor = options.name;
	
	
	// item.append(optionsLabel);
	// item.append(options);

	//var item = dynamicMcisEditor(step, isReadonly, editorContext)

	root.appendChild(item)
}

// // IF 는 이름 변경, true조건, false 조건
// // branches 아래에 true:[], false[] 가 들어있고 properties: {} 가 있음.
function appendIfbox(root, step, isReadonly, editorContext){
	const item = document.createElement('div');

	console.log("appendIfbox")
	console.log(step)

	// true
	const trueInput = document.createElement('input');
	trueInput.type = 'text';
	trueInput.name = 'trueParam';
	//trueInput.value = step.properties.provider;
	if (isReadonly) {
		trueInput.setAttribute('disabled', 'disabled');
	}

	const trueLabel = document.createElement('label');
	trueLabel.textContent = "True";
	trueLabel.htmlFor = trueInput.name;

	item.append(trueLabel);
	item.append(trueInput);

	item.append(document.createElement('br'));
	
	// false
	const falseInput = document.createElement('input');
	falseInput.type = 'text';
	falseInput.name = 'trueParam';
	//trueInput.value = step.properties.provider;

	const falseLabel = document.createElement('label');
	falseLabel.textContent = "False";
	falseLabel.htmlFor = falseInput.name;

	item.append(falseLabel);
	item.append(falseInput);


	root.appendChild(item)
}

// // branch를 삭제할 수 있는 delete branch
// // 새로운 branch를 추가할 수 있는 add branch 
function appendParallelbox(root, step, isReadonly, editorContext){

	function deleteBranch(branch, name) {
		root.removeChild(branch);
		delete step.branches[name];
		editorContext.notifyChildrenChanged();
	}

	// parallel 안에서 branch 추가
	function appendBranch(name) {
		const newBranch = document.createElement('div');
		newBranch.className = 'switch-branch';

		const newBranchTitle = document.createElement('h4');
		newBranchTitle.innerText = name;

		const newBranchLabel = document.createElement('label');
		newBranchLabel.innerText = 'Condition: ';

		newBranch.appendChild(newBranchTitle);
		newBranch.appendChild(newBranchLabel);
		console.log("addBranch " , step.properties)
		appendConditionEditor(newBranch, step.properties.conditions[name], value => {
			step.properties.conditions[name] = value;
			editorContext.notifyPropertiesChanged();
		});

		const deleteButton = createButton('Delete branch', () => {
			if (Object.keys(step.branches).length <= 1) {
				window.alert('You cannot delete the last branch.');
				return;
			}
			if (!window.confirm('Are you sure?')) {
				return;
			}
			deleteBranch(newBranch, name);
		});
		newBranch.appendChild(deleteButton);
		root.appendChild(newBranch);
	}

	function addBranch(name) {
		step.properties.conditions[name] = "add condition";
		step.branches[name] = [];
		editorContext.notifyChildrenChanged();
		appendBranch(name);
	}

	const item = document.createElement('div');

	if (!isReadonly) {// 수정mode일 때, branch 추가버튼		
		const addBranchButton = createButton('Add branch', () => {
			const branchName = window.prompt('Enter branch name');
			if (branchName) {
				addBranch(branchName);
			}
		});	

		appendPropertyTitle(root, 'Branches');
		root.appendChild(addBranchButton);
	}

	for (const initialBranchName of Object.keys(step.branches)) {
		appendBranch(initialBranchName);
	}

	console.log("appendParallelbox")
	console.log(step)
}

// Step 수정 시
function appendStepEdit(root, step, isReadonly, editorContext) {
	console.log("appendStepEdit ", step)
	console.log("appendStepEdit editorContext ", editorContext)
	// title set & change 설정
	appendTitlebox(root, step, isReadonly, editorContext);

	// Task Group은 property 수정이 없음.
	if (step.componentType === 'container' && step.type === 'loop') {
		return;
	}

	// 기본 task component( task, write, if, parallel ) 영역 ////////////////////////////////
	if (step.componentType === 'task' && step.type === 'task') {
		appendTaskPropertiesbox(root, step, isReadonly, editorContext);
		return;
	}

	if (step.componentType === 'task' && step.type === 'write') {
		appendTaskPropertiesbox(root, step, isReadonly, editorContext);
		return;
	}

	if (step.componentType === 'switch' && step.type === 'if') {
		appendIfbox(root, step, isReadonly, editorContext);
		return;
	}

	if (step.componentType === 'switch' && step.type === 'parallel') {
		appendParallelbox(root, step, isReadonly, editorContext);
		return;
	}

    /////////// 추가 task component 영역 //////////////////////////////
    if( step.componentType === 'task' && step.type == "McisDynamic"){
		//dynamicMcisEditor(step, isReadonly, editorContext);
        appendDynamicMcisTaskPropertiesbox(root, step, isReadonly, editorContext)
        return;
	}

}

// Step: task의 Title 표시 
function appendTitlebox(root, step, isReadonly, editorContext) {	
	const item = document.createElement('div');	

	let titleLabel = ""
	let titleText = step.name
	if (step.componentType === 'container' && step.type === 'loop') {
		titleLabel = "Task Group"		
	}else if (step.componentType === 'task'){
		titleLabel = "Task"
	}else if (step.componentType === 'switch' && step.type === 'if'){
		titleLabel = "IF"
	}else if (step.componentType === 'switch' && step.type === 'parallel'){
		titleLabel = "Parallel"
	}else{
		titleLabel = step.componentType + " : " + step.type
	}

	const childTitle = document.createElement('h2');
	childTitle.innerText = titleLabel + " Component";

	const childParamInput = document.createElement('input');
	childParamInput.type = 'text';
	childParamInput.name = 'titleParam';
	childParamInput.value = titleText;

	//if (isReadonly) {
		childParamInput.setAttribute('disabled', 'disabled');// task 명은 수정불가(?) : 사용자가 입력하도록 해야하나? attribute를 별도로 가져가야겠지?
	// }else{
	// 	// 수정 모드일 때 event listener 설정
	// 	childParamInput.addEventListener('input', () => {
	// 		console.log("eventListener ", childParamInput.value)
	// 			step.name = childParamInput.value;
	// 			editorContext.notifyNameChanged();
	// 	});
	// }	

	item.append(childTitle);
	item.append(childParamInput);

	root.appendChild(item);
}

// Step : 현재 step까지의 path 표시.
function appendPath(root, parents, step) {
	// console.log("appendPath")
	// console.log(root)
	// console.log(step)	
	const path = document.createElement('div');
	path.className = 'step-path';
	path.innerText = 'Step path: ' + parents.map((parent) => {
		console.log("parent ", parent)
		return typeof parent === 'string' ? parent : parent.name;
	}).join('/');
	root.appendChild(path);
}

////////////////////////////// custom task component //////////////////////////////////////
// createXxxStep
// appendXxxBox
// step 안에서  appendTitlebox, appendPath

// custom task component
function defineTaskStepDynamicMcis(id, name, properties) {
	console.log("createTaskStepDynamicMcis id", id)
	console.log("createTaskStepDynamicMcis name", name)
	console.log("createTaskStepDynamicMcis ", properties)
	
	return {
		id,
		componentType: 'task',
		type: 'McisDynamic', // Toolbox 에 표시되는 이름
		name: name, // task의 이름
		properties: properties || {
			provider: 'defaultProvider',
			options: '{}',
            body: request_body
		}
	};
}
// 기존 예제의 function 명이 createTaskXXX여서 남겨 둠.
// function createTaskStepDynamicMcis(id, type, name, properties) {
// 	console.log("createTaskStepDynamicMcis id", id)
// 	console.log("createTaskStepDynamicMcis type", type)
// 	console.log("createTaskStepDynamicMcis name", name)
// 	console.log("createTaskStepDynamicMcis ", properties)
	
// 	return {
// 		id: id,
// 		componentType: 'task',
// 		type,
// 		name,
// 		properties: properties || {
// 			provider: 'defaultProvider',
// 			options: '{}'

// 		}
// 	};
// }

// DynamicMcis Task 입력 내용 표시
function appendDynamicMcisTaskPropertiesbox(root, step, isReadonly, editorContext){
    
    console.log("in appendDynamicMcisTaskPropertiesbox step ", step)
    console.log("in appendDynamicMcisTaskPropertiesbox editorContext ", editorContext)
    var stepBody = step.properties.body;
    var stepVm = step.properties.body.vm[0];
	// Editor 정의
	// Mcis는 여러개의 VM을 가질 수 있으므로 id는 순번을 붙인다.(addVM 하면서 순번 추가할 수 있음)
	const item = document.createElement('div');

    // MCIS 영력
    const mcisItem = document.createElement('div');
    var addMcis = "";
    addMcis += '<ul>';
	
	addMcis += '<li>';
	addMcis += '	<label>MCIS Label</label>';
	addMcis += '	<input type="text" name="mcisLabel" value="' + stepBody.label + '" title="" id="mcisLabel" />';
	addMcis += '</li>';
    addMcis += '</ul>';
    mcisItem.innerHTML = addMcis;
    item.appendChild(mcisItem);

    const hr = document.createElement('hr');
    item.appendChild(hr);

    // VM 영역
    const vmItem = document.createElement('div');
	var addServer = "";
	addServer += '<ul>';
	
	addServer += '<li>';
	addServer += '	<label>VM Label</label>';
	addServer += '	<input type="text" name="vmLabel" value="' + stepVm.label + '" title="" id="vmLabel0" />';
	addServer += '</li>';

	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>VM Name</label>';
	addServer += '	<input type="text" name="vmName" value="' + stepVm.name + '" title="" id="vmName0" />';
	addServer += '</li>';

	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>VM Size</label>';
	addServer += '	<input type="text" name="subGroupSize" value="' + stepVm.subGroupSize + '" title="" id="subgroupSize0" />';
	addServer += '</li>';

	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>Spec</label>';
	addServer += '	<input type="text" name="vmSpec" value="' + stepVm.commonSpec + '" title="" id="vmSpec0" />';
	addServer += '</li>';

	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>OS Image</label>';
	addServer += '	<input type="text" name="vmImage" value="' + stepVm.commonImage + '" title="" id="vmImage0" />';
	addServer += '</li>';


	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>Root Disk Type</label>';
	addServer += '	<input type="text" name="rootDiskType" value="' + stepVm.rootDiskType + '" title="" id="rootDiskType0" />';
	addServer += '</li>';

	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>Root Disk Size</label>';
	addServer += '	<input type="text" name="rootDiskSize" value="' + stepVm.rootDiskSize + '" title="" id="rootDiskSize0" />';
	addServer += '</li>';

	addServer += '<li>';
	addServer += '	<label><span class="ch">*</span>Connection</label>';
	addServer += '	<input type="text" name="connectionName" value="' + stepVm.connectionName + '" title="" id="connectionName0" />';
	addServer += '</li>';

	addServer += '</ul>';
    vmItem.innerHTML = addServer;
    item.appendChild(vmItem);
	//item.innerHTML = addServer;
    //item.appendChild(addServer);
	root.appendChild(item)

    // event 정의 추가
    var inputMcisLabel = item.querySelector('input[name="mcisLabel"]');
    inputMcisLabel.addEventListener('change', function(event) {            
        step.properties.body.label = inputMcisLabel.value;            
        editorContext.notifyNameChanged();
        console.log('mcis label 값이 변경되었습니다.', step);
    });
    
    if( !isReadonly){
        var inputVmLabel = item.querySelector('input[name="vmLabel"]');
        inputVmLabel.addEventListener('change', function(event) {            
            step.properties.body.vm[0].label = inputVmLabel.value;            
            editorContext.notifyNameChanged();
            console.log('vm label 값이 변경되었습니다.', step);
        });

        var inputVmName = item.querySelector('input[name="vmName"]');
        inputVmName.addEventListener('change', function(event) {            
            step.properties.body.vm[0].name = inputVmName.value;            
            editorContext.notifyNameChanged();
            console.log('vm Name 값이 변경되었습니다.', step);
        });

        var inputSubGroupSize = item.querySelector('input[name="subGroupSize"]');
        inputSubGroupSize.addEventListener('change', function(event) {            
            step.properties.body.vm[0].subGroupSize = inputSubGroupSize.value;            
            editorContext.notifyNameChanged();
            console.log('vm size 값이 변경되었습니다.', step);
        });

        var inputVmSpec = item.querySelector('input[name="vmSpec"]');
        inputVmSpec.addEventListener('change', function(event) {            
            step.properties.body.vm[0].commonSpec = inputVmSpec.value;            
            editorContext.notifyNameChanged();
            console.log('vm spec 값이 변경되었습니다.', step);
        });

        var inputVmImage = item.querySelector('input[name="vmImage"]');
        inputVmImage.addEventListener('change', function(event) {            
            step.properties.body.vm[0].commonImage = inputVmImage.value;            
            editorContext.notifyNameChanged();
            console.log('vm Image 값이 변경되었습니다.', step);
        });

        var inputRootDiskType = item.querySelector('input[name="rootDiskType"]');
        inputRootDiskType.addEventListener('change', function(event) {            
            step.properties.body.vm[0].rootDiskType = inputRootDiskType.value;            
            editorContext.notifyNameChanged();
            console.log('rootDiskType 값이 변경되었습니다.', step);
        });
        var inputRootDiskSize = item.querySelector('input[name="rootDiskSize"]');
        inputRootDiskSize.addEventListener('change', function(event) {            
            step.properties.body.vm[0].rootDiskSize = inputRootDiskSize.value;            
            editorContext.notifyNameChanged();
            console.log('rootDiskSize 값이 변경되었습니다.', step);
        });
        var inputConnectionName = item.querySelector('input[name="connectionName"]');
        inputConnectionName.addEventListener('change', function(event) {            
            step.properties.body.vm[0].connectionName = inputConnectionName.value;            
            editorContext.notifyNameChanged();
            console.log('connectionName 값이 변경되었습니다.', step);
        });
    }
    return
}


function addVM() {
    // 이 부분에 addVM 함수의 내용을 작성
    console.log('addVM 함수를 호출했습니다.');
}



////// 전체 workflow에 대한 내용 확인. 
// appendTitle, appendTextField 정의
function customRootEditorProvider(definition, editorContext, isReadonly){
    const rootEditorContainer = document.createElememt('span');
    appendTitlebox(rootEditorContainer, 'Workflow overview');
    appendTextField(rootEditorContainer, 'Speed(ms)', isReadonly, definition.properties['speed'], v => {
        definition.properties['speed'] = parseInt(v,10);
        editorContext.notifyPropertiesChanged();
    });
    return rootEditorContainer
}