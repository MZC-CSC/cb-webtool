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
	definition: getStartDefinition()	
}

// workflowId로 조회
function loadState(workflowId) {
    alert("get workflow ", workflowId)
	// const state = localStorage[localStorageKey];
	// if (state) {
	// 	return JSON.parse(state);
	// }
	// return {
	// 	definition: getStartDefinition()
	// }
}

// wor
const configuration = {
	undoStackSize: 20,
	undoStack: initialState.undoStack,

	toolbox: {
		groups: [
			toolboxGroup('Main'),
			toolboxGroup('File system'),
			toolboxGroup('E-mail')
		]
	},

	controlBar: true,

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

	editors: {
		rootEditorProvider: (definition, _context, isReadonly) => {
			const root = document.createElement('div');
			root.className = 'definition-json';
			root.innerHTML = '<textarea style="width: 100%; border: 0;" rows="50"></textarea>';
			const textarea = root.getElementsByTagName('textarea')[0];
			if (isReadonly) {
				textarea.setAttribute('readonly', 'readonly');
			}
			textarea.value = JSON.stringify(definition, null, 2);
			return root;
		},

		stepEditorProvider: (step, editorContext, _definition, isReadonly) => {
			const root = document.createElement('div');

			appendCheckbox(root, 'Is invalid', isReadonly, !!step.properties['isInvalid'], (checked) => {
				step.properties['isInvalid'] = checked;
				editorContext.notifyPropertiesChanged();
			});

			if (step.type === 'if') {
				appendCheckbox(root, 'Catch branch', isReadonly, !!step.branches['catch'], (checked) => {
					if (checked) {
						step.branches['catch'] = [];
					} else {
						delete step.branches['catch'];
					}
					editorContext.notifyChildrenChanged();
				});
			}

			appendPath(root, step);
			return root;
		}
	}
};

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
//const placeholder = document.getElementById('designer');
document.addEventListener('DOMContentLoaded', function() {
    
    console.log("DOMContentLoaded ")

    const definition = {
    properties: {
        'myProperty': 'my-value',
        // root properties...
    },
    sequence: [
        // steps...
    ]
    };
    console.log("workflow definition ")

    // const configuration = {
    // theme: 'light', // optional, default: 'light'
    // isReadonly: false, // optional, default: false
    // undoStackSize: 10, // optional, default: 0 - disabled, 1+ - enabled

    // steps: {
    //     // all properties in this section are optional

    //     iconUrlProvider: (componentType, type) => {
    //     return `icon-${componentType}-${type}.svg`;
    //     },

    //     isDraggable: (step, parentSequence) => {
    //     return step.name !== 'y';
    //     },
    //     isDeletable: (step, parentSequence) => {
    //     return step.properties['isDeletable'];
    //     },
    //     isDuplicable: (step, parentSequence) => {
    //         return true;
    //     },
    //     canInsertStep: (step, targetSequence, targetIndex) => {
    //     return targetSequence.length < 5;
    //     },
    //     canMoveStep: (sourceSequence, step, targetSequence, targetIndex) => {
    //     return !step.properties['isLocked'];
    //     },
    //     canDeleteStep: (step, parentSequence) => {
    //     return step.name !== 'x';
    //     }
    // },

    // validator: {
    //     // all validators are optional

    //     step: (step, parentSequence, definition) => {
    //     return /^[a-z]+$/.test(step.name);
    //     },
    //     root: (definition) => {
    //     return definition.properties['memory'] > 256;
    //     }
    // },

    // toolbox: {
    //     isCollapsed: false,
    //     groups: [
    //     {
    //         name: 'Files',
    //         steps: [
    //         // steps for the toolbox's group
    //         ]
    //     },
    //     {
    //         name: 'Notification',
    //         steps: [
    //         // steps for the toolbox's group
    //         ]
    //     }
    //     ]
    // },

    // editors: {
    //     isCollapsed: false,
    //     rootEditorProvider: (definition, rootContext, isReadonly) => {
    //     const editor = document.createElement('div');
    //     // ...
    //     return editor;
    //     },
    //     stepEditorProvider: (step, stepContext, definition, isReadonly) => {
    //     const editor = document.createElement('div');
    //     // ...
    //     return editor;
    //     }
    // },

    // controlBar: true,
    // contextMenu: true,
    // };
    console.log("workflow configuration ")
    const designer = sequentialWorkflowDesigner.Designer.create(placeholder, definition, configuration);
    designer.onDefinitionChanged.subscribe((newDefinition) => {
        console.log("workflow designer.onDefinitionChanged ")
    // ...
    });

    console.log("workflow designer created ")
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