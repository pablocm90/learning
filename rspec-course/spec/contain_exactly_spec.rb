RSpec.describe 'contain exactly' do
  subject { [1, 2, 3]}

  it 'should check for the presence of all elements' do 
    expect(subject).to contain_exactly(2, 3, 1)
  end

  it {is_expected.to contain_exactly(2,1,3)}
end
