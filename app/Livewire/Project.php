<?php

namespace App\Livewire;

use Livewire\Component;
use App\Models\Project as Projects;

class Project extends Component
{
    public int $id;
    public string $name, $description;
    public bool $updating = false;

    /**
     * @var Projects[] $projects
     */
    public $projects;

    protected $listeners = [
        'editProject' => 'edit',
        'deleteProject' => 'delete',
    ];

    protected $rules = [
        'name' => 'required|string|max:32',
        'description' => 'required|string|max:4096',
    ];

    public function render()
    {
        $this->projects = Projects::all();

        return view('livewire.project');
    }

    public function resetFields()
    {
        $this->name = '';
        $this->description = '';
    }

    public function store()
    {
        $this->validate();
        try {
            Projects::create([
                'name' => $this->name,
                'description' => $this->description
            ]);

            // TODO: migrate messages to a dictionary of language files for better i18n
            session()->flash('success', 'Project created successfully!');
            $this->resetFields();
        } catch (\Exception $e) {
            session()->flash('error', 'Something goes wrong while creating project!!');
            $this->resetFields();
        }
    }

    public function edit($id)
    {
        $category = Projects::findOrFail($id);
        $this->name = $category->name;
        $this->description = $category->description;
        $this->id = $category->id;
        $this->updating = true;
    }

    public function cancel()
    {
        $this->updating = false;
        $this->resetFields();
    }

    public function update()
    {
        $this->validate();
        try {
            Projects::find($this->id)->fill([
                'name' => $this->name,
                'description' => $this->description
            ])->save();

            session()->flash('success', 'Category Updated Successfully!!');

            $this->cancel();
        } catch (\Exception $e) {
            session()->flash('error', 'Something goes wrong while updating category!!');
            $this->cancel();
        }
    }

    public function destroy($id)
    {
        try {
            Projects::find($id)->delete();
            session()->flash('success', "Category Deleted Successfully!!");
        } catch (\Exception $e) {
            session()->flash('error', "Something goes wrong while deleting category!!");
        }
    }
}
