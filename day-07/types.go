package dayseven

type FileSystemItem interface {
	Size() int
	Name() string
}

type File struct {
	name string
	size int
}
func (f File) Size() int {
	return f.size
}
func (f File) Name() string {
	return f.name
}
func newFile(name string, size int) *File {
	return &File{name, size}
}

type Directory struct {
	name        string
	files       []*File
	directories []*Directory
	parent			*Directory
}
func (d Directory) FindDir(name string) *Directory {
	for _, dir := range d.directories {
		if dir.Name() == name {
			return dir
		}
	}

	return nil
}
func (d Directory) Size() int {
	size := 0
	for _, f := range d.files {
		size += f.Size()
	}
	for _, dir := range d.directories {
		size += dir.Size()
	}
	return size
}
func (d Directory) Name() string {
	return d.name
}
func newDirectory(name string) *Directory {
	d := &Directory{}
	d.name = name
	return d
}
func (d *Directory) makeChildDirectory(name string) *Directory {
	newD := newDirectory(name)
	d.directories = append(d.directories, newD)
	newD.parent = d
	return newD
}
