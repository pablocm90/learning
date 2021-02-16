RSpec.describe Hash do 
  subject(:bob) do
    { a: 1, b: 2 }
  end

  it 'has two key value pairs' do
    expect(subject.length).to eq(2)
    expect(bob.length).to eq(2)
  end
end


RSpec.describe Array do
  subject(:sally) do 
    [2, 1]
  end
  it 'has 2 elements' do 
    expect(subject.length).to eq(2)
    subject.pop
    expect(subject.length).to eq(1)
  end
end
