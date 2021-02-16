RSpec.describe 'be_matcher' do
  it 'can test for truthiness' do
    expect(true).to be_truthy
    expect('Hello').to be_truthy
  end

  it 'can test for falsyness' do
    expect(nil).to be_falsy
  end

  it 'can test for nil' do 
    expect(nil).to be_nil
  end
end
