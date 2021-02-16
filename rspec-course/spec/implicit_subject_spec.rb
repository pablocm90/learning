RSpec.describe Hash do
  it 'should start off empty' do 
    expect(subject.length).to eq(0)
  end
end


RSpec.describe Array do 
  it 'should start off empty' do 
    expect(subject.length).to eq(0)
    subject.push('A')
    expect(subject.length).to eq(1)
  end
end