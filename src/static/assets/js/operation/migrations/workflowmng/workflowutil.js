//////////////////////// Canvas 정의 영역 ////////////////////
function defaultState() {
	return {definition: getDefaultDefinition()}
    // return {
    //     definition: getStartDefinition()
    // }
}

// workflow 기본 정의 : 비어있는 workflow
function getDefaultDefinition() {
	return {
		properties: {
            'workflow': 'myworkflow',
        },
		sequence: [	]
	};
}

// sample workflow
function getStartDefinition() {
	return {
		properties: {},
		sequence: [
			createIfStep('00000000000000000000000000000001',
				[ createTaskStep('00000000000000000000000000000002', 'save', 'Save file', { isInvalid: true }) ],
				[ createTaskStep('00000000000000000000000000000003', 'text', 'Send email') ]
			),
			createContainerStep('00000000000000000000000000000004', [
				createTaskStep('00000000000000000000000000000005', 'task', 'Create task')
			])
		]
	};
}
//////////////////////// Canvas 정의 영역 ////////////////////






















// // editor와 view에서 사용할 function들 

// // Task 생성 
// //createTaskStep('task', 'task1')
// function createTaskStep(id, type, name, properties) {
// 	console.log("createTaskStep ", properties)
// 	return {
// 		id: id,
// 		componentType: 'task',
// 		type,
// 		name,
// 		properties: properties || ''
// 	};
// }

// // Airflow Taks는 provider + options 로 되어 있어 대응하는 Task
// // createAirflowTaskStep(null, 'write', 'task'),
// function createAirflowTaskStep(id, type, name, properties) {
// 	console.log("createAirflowTaskStep ", properties)
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

// // TODO : 속성을 읽어서 TASK에 필요한 항목으로 정의.( name, json 으로 할까?)
// // 지정형식이 아닌 형태의 TASK를 모델로 추가.
// // 사용은 어떻게?? // getToolboxGroup 에서 사용할 task model을 create 함.
// // task에 필요한 속성정의
// function createAnyTaskStep(id, name, properties){
// 	console.log("createAnyTaskStep ", properties)
// }

// // container 로 감싸는 것으로 포함할 step들을 넘김. 빈값도 가능
// function createContainerStep(id, containerName, steps) {
// 	console.log("createContainerStep")
// 	return {
// 		id: id,
// 		componentType: 'container',
// 		type: 'loop',
// 		name: containerName || 'Task Group',
// 		properties: {},
// 		sequence: steps
// 	};
// }

// // if step : truestep과 falsestep을 인자로 받음. 빈값도 가능,  세가지 조건(true, false, extra) 가능
// function createIfStep(id, trueSteps, falseSteps, extraSteps) {
// 	console.log("createIfStep")

// 	// 기본은 true, false 두개임. 
// 	let childBranches = {
// 		'true': trueSteps || [],
// 		'false': falseSteps || [],
// 	}

// 	if( extraSteps){
// 		childBranches['extra'] = extraSteps;
// 	}

// 	return {
// 		id: id,
// 		componentType: 'switch',
// 		type: 'if',
// 		name: 'If',
// 		branches: childBranches,
// 		properties: {}
// 	};
// }

// // parallel step : 병렬 step으로 children 만큼 병렬로 진행
// // child branch와 child condition은 쌍으로 존재해야 함. 맞지 않으면 수정할 때 에러 남.
// function createParallelStep(id, name, children) {
// 	console.log("createParallelStep")
// 	let childrenBranches = {}
// 	let childrenConditions = {}

// 	// branches는 최소 1개 필요.
// 	if(children && children.length > 0){
// 		for (let i = 0; i < children.length; i++) {
// 			childrenBranches[`Condition ${String.fromCharCode(65 + i)}`] = [children[i]];
// 			childrenConditions[`Condition ${String.fromCharCode(65 + i)}`] = [children[i].name];
// 		}
// 	}else{// branch는 기본 3개.
// 		childrenBranches[`Condition1`] = [];
// 		childrenBranches[`Condition2`] = [];
// 		childrenBranches[`Condition3`] = [];
		
// 		childrenConditions[`Condition1`] = 'Condition1Value';
// 		childrenConditions[`Condition2`] = 'Condition2Value';
// 		childrenConditions[`Condition3`] = 'Condition3Value';
// 	}
		
// 	const parallel = {
// 		id: id,
// 		componentType: 'switch',
// 		type: 'parallel',		
// 		name,
// 		properties: {
// 			conditions: childrenConditions,
// 		},
// 		// properties: {
// 			//conditions: childrenConditions
// 			// "Condition1": "1",
// 			// "Condition2": "2",
// 			// "Condition3": "3",
// 		// },
// 		branches: childrenBranches
// 		// branches: {
// 		// 	"Condition1": [],
// 		// 	// "Condition2": [],
// 		// 	// "Condition3": [],
// 		// }
// 	}
// 	//console.log(parallel)
// 	return parallel;
// }

// // Readonly 설정 표시 : editor는 설정변경하지 않음.
// // function reloadChangeReadonlyButtonText() {
// // 	changeReadonlyButton.innerText = 'Readonly: ' + (designer.isReadonly() ? 'ON' : 'OFF');
// // }

// // 상세창에 checkbox 표시
// function appendCheckbox(root, label, isReadonly, isChecked, onClick) {
// 	const item = document.createElement('div');
// 	item.innerHTML = '<div><h3></h3> <input type="checkbox" /></div>';
// 	const h3 = item.getElementsByTagName('h3')[0];
// 	h3.innerText = label;
// 	const input = item.getElementsByTagName('input')[0];
// 	input.checked = isChecked;
// 	if (isReadonly) {
// 		input.setAttribute('disabled', 'disabled');
// 	}
// 	input.addEventListener('click', () => {
// 		onClick(input.checked);
// 	});
// 	root.appendChild(item);
// }

// // 상세 창에 선택된 항목의 title 표시 및 event set
// // isReadonly=false일 때 변경과 동시에 바로 반영
// function appendTitlebox(root, step, isReadonly, editorContext) {	
// 	const item = document.createElement('div');	

// 	let titleLabel = ""
// 	let titleText = step.name
// 	if (step.componentType === 'container' && step.type === 'loop') {
// 		titleLabel = "Task Group"		
// 	}else if (step.componentType === 'task'){
// 		titleLabel = "Task"
// 	}else if (step.componentType === 'switch' && step.type === 'if'){
// 		titleLabel = "IF"
// 	}else if (step.componentType === 'switch' && step.type === 'parallel'){
// 		titleLabel = "Parallel"
// 	}else{
// 		titleLabel = step.componentType + " : " + step.type
// 	}

// 	const childTitle = document.createElement('h2');
// 	childTitle.innerText = titleLabel;

// 	const childParamInput = document.createElement('input');
// 	childParamInput.type = 'text';
// 	childParamInput.name = 'titleParam';
// 	childParamInput.value = titleText;

// 	if (isReadonly) {
// 		childParamInput.setAttribute('disabled', 'disabled');
// 	}else{
// 		// 수정 모드일 때 event listener 설정
// 		childParamInput.addEventListener('input', () => {
// 			console.log("eventListener ", childParamInput.value)
// 				step.name = childParamInput.value;
// 				editorContext.notifyNameChanged();
// 		});
// 	}	

// 	item.append(childTitle);
// 	item.append(childParamInput);

// 	root.appendChild(item);
// }

// // TASK Group 은 이름만 변경.
// // TASK 는 이름 변경, properties 변경
// // IF 는 이름 변경, true조건, false 조건
// // ...
// function appendTaskPropertiesbox(root, step, isReadonly, editorContext){

// 	const item = document.createElement('div');

// 	console.log("appendTaskPropertiesbox")
// 	console.log(step)

// 	// 기본 상항으로 사용할 provider와 option

// 	// provider : input
// 	const providerInput = document.createElement('input');
// 	providerInput.type = 'text';
// 	providerInput.name = 'inputParam';
// 	providerInput.value = step.properties.provider;
// 	if (isReadonly) {
// 		providerInput.setAttribute('disabled', 'disabled');
// 	}

// 	const providerLabel = document.createElement('label');
// 	providerLabel.textContent = "Provider";
// 	providerLabel.htmlFor = providerInput.name;
		
// 	item.append(providerLabel);
// 	item.append(providerInput);

// 	// options : text area => json format
// 	const options = document.createElement('textarea');
// 	options.name = 'optionParams';
// 	options.rows = 4;
// 	options.cols = 10
// 	options.value = step.properties.options;
// 	if (isReadonly) {
// 		options.setAttribute('disabled', 'disabled');
// 	}

// 	const optionsLabel = document.createElement('label');
// 	optionsLabel.textContent = "Options";
// 	optionsLabel.htmlFor = options.name;
	
	
// 	item.append(optionsLabel);
// 	item.append(options);


// 	root.appendChild(item)
// }

// // IF 는 이름 변경, true조건, false 조건
// // branches 아래에 true:[], false[] 가 들어있고 properties: {} 가 있음.
// function appendIfbox(root, step, isReadonly, editorContext){
// 	const item = document.createElement('div');

// 	console.log("appendIfbox")
// 	console.log(step)

// 	// true
// 	const trueInput = document.createElement('input');
// 	trueInput.type = 'text';
// 	trueInput.name = 'trueParam';
// 	//trueInput.value = step.properties.provider;
// 	if (isReadonly) {
// 		trueInput.setAttribute('disabled', 'disabled');
// 	}

// 	const trueLabel = document.createElement('label');
// 	trueLabel.textContent = "True";
// 	trueLabel.htmlFor = trueInput.name;

// 	item.append(trueLabel);
// 	item.append(trueInput);

// 	item.append(document.createElement('br'));
	
// 	// false
// 	const falseInput = document.createElement('input');
// 	falseInput.type = 'text';
// 	falseInput.name = 'trueParam';
// 	//trueInput.value = step.properties.provider;

// 	const falseLabel = document.createElement('label');
// 	falseLabel.textContent = "False";
// 	falseLabel.htmlFor = falseInput.name;

// 	item.append(falseLabel);
// 	item.append(falseInput);


// 	root.appendChild(item)
// }

// // branch를 삭제할 수 있는 delete branch
// // 새로운 branch를 추가할 수 있는 add branch 
// function appendParallelbox(root, step, isReadonly, editorContext){

// 	function deleteBranch(branch, name) {
// 		root.removeChild(branch);
// 		delete step.branches[name];
// 		editorContext.notifyChildrenChanged();
// 	}

// 	// parallel 안에서 branch 추가
// 	function appendBranch(name) {
// 		const newBranch = document.createElement('div');
// 		newBranch.className = 'switch-branch';

// 		const newBranchTitle = document.createElement('h4');
// 		newBranchTitle.innerText = name;

// 		const newBranchLabel = document.createElement('label');
// 		newBranchLabel.innerText = 'Condition: ';

// 		newBranch.appendChild(newBranchTitle);
// 		newBranch.appendChild(newBranchLabel);
// 		console.log("addBranch " , step.properties)
// 		appendConditionEditor(newBranch, step.properties.conditions[name], value => {
// 			step.properties.conditions[name] = value;
// 			editorContext.notifyPropertiesChanged();
// 		});

// 		const deleteButton = createButton('Delete branch', () => {
// 			if (Object.keys(step.branches).length <= 1) {
// 				window.alert('You cannot delete the last branch.');
// 				return;
// 			}
// 			if (!window.confirm('Are you sure?')) {
// 				return;
// 			}
// 			deleteBranch(newBranch, name);
// 		});
// 		newBranch.appendChild(deleteButton);
// 		root.appendChild(newBranch);
// 	}

// 	function addBranch(name) {
// 		step.properties.conditions[name] = "add condition";
// 		step.branches[name] = [];
// 		editorContext.notifyChildrenChanged();
// 		appendBranch(name);
// 	}

// 	const item = document.createElement('div');

// 	if (!isReadonly) {// 수정mode일 때, branch 추가버튼		
// 		const addBranchButton = createButton('Add branch', () => {
// 			const branchName = window.prompt('Enter branch name');
// 			if (branchName) {
// 				addBranch(branchName);
// 			}
// 		});	

// 		appendPropertyTitle(root, 'Branches');
// 		root.appendChild(addBranchButton);
// 	}

// 	for (const initialBranchName of Object.keys(step.branches)) {
// 		appendBranch(initialBranchName);
// 	}

// 	console.log("appendParallelbox")
// 	console.log(step)
// }

// // 상세 창에 inputbox 표시
// function appendInputbox(root, label, isReadonly, textValue) {
// 	const item = document.createElement('div');	
	
// 	const childTitle = document.createElement('h2');
// 	//childTitle.innerText = label;

// 	const childParamInput = document.createElement('input');
// 	childParamInput.type = 'text';
// 	childParamInput.name = 'inputParam';
// 	childParamInput.value = textValue;
// 	if (isReadonly) {
// 		childParamInput.setAttribute('disabled', 'disabled');
// 	}

// 	const childParamLabel = document.createElement('label');
// 	childParamLabel.textContent = label;
// 	childParamLabel.htmlFor = childParamInput.name;
	
// 	item.append(childTitle);
// 	item.append(childParamLabel);
// 	item.append(childParamInput);

// 	root.appendChild(item);
// }

// // parallel 에서 사용
// function createButton(text, clickHandler) {
// 	const button = document.createElement('button');
// 	button.innerText = text;
// 	button.addEventListener('click', clickHandler, false);
// 	return button;
// }

// // parallel 에서 사용
// function appendPropertyTitle(root, title) {
// 	const h3 = document.createElement('h3');
// 	h3.innerText = title;
// 	root.appendChild(h3);
// }

// // parallel 에서 사용
// function appendConditionEditor(root, value, onChange) {
// 	const input = document.createElement('input');
// 	input.type = 'text';
// 	input.value = value;
// 	input.addEventListener(
// 		'input',
// 		() => {
// 			onChange(input.value);
// 		},
// 		false
// 	);
// 	root.appendChild(input);
// }

// ///////////// end of private functions //////////////////

// ///////////// export functions //////////////////////////
// // ToolBox 표시할 때 Group 정의
// // toolbox : category별로 group을 다르게 설정
// export function getToolboxGroup(category) {
// 	console.log("getToolBox ", category)
// 	switch (category) {
// 		case "logic" :
// 			return {
// 				name: 'logic',
// 				steps: [
// 					createIfStep(null,[], []),
// 					createParallelStep(null, "parallels", []),					
// 				]
// 			};
// 		case "step" :
// 			return {
// 				name: 'step',
// 				steps: [
// 					createContainerStep(null,'task group', []),					
// 					createAirflowTaskStep(null, 'write', 'task'),
// 				]
// 			};
// 	}	
// }


// // STEP 편집 
// // 편집한 내용을 반영하려면 editorContext에 notify...() 를 호출해야한다.
// // step의 이름 변경 : notifyNameChanged()
// // step내 property 변경 : notifyPropertiesChanged()
// export function appendStepEdit(root, step, isReadonly, editorContext) {
// 	console.log("appendStepEdit ", step)
// 	// title set & change 설정
// 	appendTitlebox(root, step, isReadonly, editorContext);

// 	// Task Group은 property 수정이 없음.
// 	if (step.componentType === 'container' && step.type === 'loop') {
// 		return;
// 	}

// 	// 기본 task 설정
// 	if (step.componentType === 'task' && step.type === 'task') {
// 		appendTaskPropertiesbox(root, step, isReadonly, editorContext);
//         return;
// 	}

// 	if (step.componentType === 'task' && step.type === 'write') {
// 		appendTaskPropertiesbox(root, step, isReadonly, editorContext);
//         return;
// 	}

// 	if (step.componentType === 'switch' && step.type === 'if') {
// 		appendIfbox(root, step, isReadonly, editorContext);
//         return;
// 	}

// 	if (step.componentType === 'switch' && step.type === 'parallel') {
// 		appendParallelbox(root, step, isReadonly, editorContext);
//         return;
// 	}

// }

// // 현재 step까지의 path 표시.
// export function appendPath(root, parents, step) {
// 	// console.log("appendPath")
// 	// console.log(root)
// 	// console.log(step)	
// 	const path = document.createElement('div');
// 	path.className = 'step-path';
// 	path.innerText = 'Step path: ' + parents.map((parent) => {
// 		console.log("parent ", parent)
// 		return typeof parent === 'string' ? parent : parent.name;
// 	}).join('/');
// 	root.appendChild(path);
// }

// // 기본 workflowsequence
// export function getDefaultWorkflowSequence(){
//     return createContainerStep(
//             [createTaskStep('task', 'task1')]
//         );
// }


// /*

// {
//   "properties": {},
//   "sequence": [
//     {
//       "id": "0ce635e0a81b52bcf227107202424cb1",
//       "componentType": "switch",
//       "type": "if",
//       "name": "If1a",
//       "branches": {
//         "true": [
//           {
//             "id": "e48a5bb71915602657a9bcb25c08ba1c",
//             "componentType": "task",
//             "type": "write",
//             "name": "task true",
//             "properties": {
//               "provider": "defaultProvider",
//               "options": "{}"
//             }
//           }
//         ],
//         "false": [
//           {
//             "id": "a70f2bc432454e3b0e01be4982eb4c0e",
//             "componentType": "task",
//             "type": "write",
//             "name": "task false1",
//             "properties": {
//               "provider": "defaultProvider",
//               "options": "{}"
//             }
//           }
//         ]
//       },
//       "properties": {}
//     },
// */
// // root에서 step id로 조회
// export function findStepById(stepId, rootJson) {
// 	const element = rootJson.sequence.find(item => item.id === stepId);
	
// 	// 요소를 찾지 못한 경우 null을 반환
// 	if (!element) {
// 	  return null;
// 	}
// 	console.log(element)
// 	// 요소를 찾은 경우 해당 요소를 반환
// 	return element;
// }