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

// 워크flowId를 가지고 조회하여 data를 표시.
const placeholder = document.getElementById('workflowplaceholder');
const localStorageKey = 'sqdFullscreen';

const initialState = defaultState();
function defaultState() {
	return {definition: getDefaultDefinition()}
    // return {
    //     definition: getStartDefinition()
    // }
}

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
}

// workflow 저장
function saveWorkflow() {
	localStorage[localStorageKey] = JSON.stringify({
		definition: designer.getDefinition(),
		undoStack: designer.dumpUndoStack()
	});
}

// workflowmng 에서는 readonly만.
const readonlyConfiguration = {
	isReadonly: true,
    
	//toolbox: {}, // toolbox가 false가 아닌 경우에는 비어있으면 안됨.
    toolbox: false,
    //contextMenu: false,
	controlBar: false,
	steps: {
		isDuplicable: () => true,
		iconUrlProvider: (_, type) => {
			return `/assets/js/sequential-workflow-designer/ico/icon-${type}.svg`
		},
	},

	validator: {
		step: (step) => {
			return !step.properties['isInvalid'];
		}
	},

    // root와 step provider를 지정하며 parameter에서 isReadonly 속성으로 구분가능.    
    editors : {
        //isCollapsed: true,
        isCollapsed: false,
        rootEditorProvider:(definition, _context, isReadonly) => {
            const root = document.createElement('div');
            root.className = 'definition-json';
			root.innerHTML = '<textarea style="width: 100%; border: 0;" rows="50"></textarea>';
			const textarea = root.getElementsByTagName('textarea')[0];
			textarea.setAttribute('readonly', 'readonly');			
			textarea.value = JSON.stringify(definition, null, 2);
            return root;
        },
        stepEditorProvider: (step, editorContext, _definition, isReadonly) => {
            const root = document.createElement('div');
            appendPath(root, step);
			return root;
        }
    },
	
};




function refreshValidationStatus() {
	validationStatusText.innerText = designer.isValid() ? 'Definition is valid' : 'Definition is invalid';
}

//const placeholder = document.getElementById('designer');
document.addEventListener('DOMContentLoaded', function() {
    
    console.log("DOMContentLoaded ")

    const definition = initialState;
    console.log("workflow definition ", definition)

    const configuration = readonlyConfiguration;

    console.log("workflow configuration ", configuration)
    console.log("placeholder ", placeholder)
    const designer = sequentialWorkflowDesigner.Designer.create(placeholder, definition.definition, configuration);
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

function appendCheckbox(root, label, isReadonly, isChecked, onClick) {
	const item = document.createElement('div');
	item.innerHTML = '<div><h3></h3> <input type="checkbox" /></div>';
	const h3 = item.getElementsByTagName('h3')[0];
	h3.innerText = label;
	const input = item.getElementsByTagName('input')[0];
	input.checked = isChecked;
	if (isReadonly) {
		input.setAttribute('disabled', 'disabled');
	}
	input.addEventListener('click', () => {
		onClick(input.checked);
	});
	root.appendChild(item);
}

function appendPath(root, step) {
	const parents = designer.getStepParents(step);
	const path = document.createElement('div');
	path.className = 'step-path';
	path.innerText = 'Step path: ' + parents.map((parent) => {
		return typeof parent === 'string' ? parent : parent.name;
	}).join('/');
	root.appendChild(path);
}